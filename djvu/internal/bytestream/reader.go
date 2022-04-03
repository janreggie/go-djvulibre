package bytestream

import (
	"errors"
	"io"
)

// Reader takes in an io.Reader but with more options
type Reader struct {
	r io.Reader
	e error
}

func NewReader(r io.Reader) Reader {
	return Reader{r: r, e: nil}
}

// ReadInteger reads the nearest integer.
// If there are errors reading the reader,
// save that error in the Reader.
func (r *Reader) ReadInteger() (uint32, error) {
	var result uint32
	br := NewByteReader(r.r)
	br.B = '\n' // some lookahead character that will be overridden anyways

	// Eat blank before integer
	for br.B == ' ' || br.B == '\t' || br.B == '\r' || br.B == '\n' || br.B == '#' {
		if br.B == '#' {
			for {
				if br.Advance(); br.E != nil {
					r.e = br.E
					return 0, br.E
				}
				if br.B == '\n' || br.B == '\r' {
					break
				}
			}
		}
		if br.Advance(); br.E != nil {
			r.e = br.E
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
		if br.Advance(); br.E != nil {
			r.e = br.E
			return 0, br.E
		}
	}

	return result, nil
}

func (r *Reader) Error() error { return r.e }

func (r *Reader) ClearError() { r.e = nil }
