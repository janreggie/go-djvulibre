package bytestream

import (
	"io"

	"github.com/pkg/errors"
)

// ByteReader allows serial reading of bytes.
// If there are errors, it will be wrapped by "could not read data".
// If there have been errors anywhere, it will not continue to read.
type ByteReader struct {
	r io.Reader
	b []byte
	B byte
	E error
}

func NewByteReader(r io.Reader) ByteReader {
	return ByteReader{
		r: r,
		b: make([]byte, 1),
		B: 0,
		E: nil,
	}
}

// Advance reads a single byte to B,
// and places an error on E
// if it does encounter one.
// If E isn't nil, it will not read a byte.
func (br *ByteReader) Advance() {
	if br.E != nil {
		return
	}

	_, br.E = br.r.Read(br.b)
	br.E = errors.Wrapf(br.E, "could not read data") // Will return nil if br.E == nil
	br.B = br.b[0]
}
