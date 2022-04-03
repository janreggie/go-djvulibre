package image

import (
	"bufio"
	"io"

	"github.com/pkg/errors"
)

func readInteger(r io.Reader) (uint32, error) {
	var result uint32
	br := newBytereader(r)
	br.B = '\n' // some lookahead character that will be overridden anyways

	// Eat blank before integer
	for br.B == ' ' || br.B == '\t' || br.B == '\r' || br.B == '\n' || br.B == '#' {
		if br.B == '#' {
			for {
				if br.advance(); br.E != nil {
					return 0, br.E
				}
				if br.B == '\n' || br.B == '\r' {
					break
				}
			}
		}
		if br.advance(); br.E != nil {
			return 0, br.E
		}
	}

	// Check integer
	if br.B < '0' || br.B > '9' {
		return 0, errors.New("could not find integer")
	}

	// Eat integer
	for br.B >= '0' && br.B <= '9' {
		result = result*10 + uint32(br.B-'0')
		if br.advance(); br.E != nil {
			return 0, br.E
		}
	}

	return result, nil
}

// bytereader allows serial reading of bytes.
// If there are errors, it will be wrapped by "could not read data".
// If there have been errors anywhere, it will not continue to read.
type bytereader struct {
	buf *bufio.Reader
	B   byte
	E   error
}

func newBytereader(r io.Reader) bytereader {
	return bytereader{
		buf: bufio.NewReader(r),
	}
}

func (br *bytereader) advance() {
	if br.E != nil {
		return
	}
	br.B, br.E = br.buf.ReadByte()
	if br.E != nil {
		br.E = errors.Wrapf(br.E, "could not read data")
	}
}
