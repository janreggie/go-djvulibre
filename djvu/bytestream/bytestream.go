package bytestream

import "io"

// Idk what to make of you so for now...
type ByteStream interface{}

// Reader is an interface which reads from a stream.
type Reader interface{ io.Reader }

// Writer is an interface which writes to a stream.
// Methods can be chained.
type Writer interface {
	io.Writer
	// TODO: Create methods like in io.Writer that "chain"

	// Write8 writes an 8-bit integer to Writer
	Write8(v uint8) Writer

	// Write16 writes a 16-bit integer to Writer
	Write16(v uint16) Writer

	// Error returns the error encountered in the process.
	// If a Writer has an Error, subsequent writes must fail.
	Error() error

	// ClearError clears whatever error there is in the Writer.
	ClearError()
}

type writer struct {
	io.Writer
	err error
}

func NewWriter(w io.Writer) Writer { return &writer{Writer: w, err: nil} }

func (w *writer) Write8(v uint8) Writer {
	if w.err != nil {
		return w
	}

	_, err := w.Write([]byte{v})
	if err != nil {
		w.err = err
	}
	return w
}

func (w *writer) Write16(v uint16) Writer {
	if w.err != nil {
		return w
	}

	bs := []byte{
		uint8((v >> 8) & 0xff),
		uint8(v & 0xff),
	}
	_, err := w.Write(bs)
	if err != nil {
		w.err = err
	}
	return w
}

func (w *writer) Error() error { return w.err }

func (w *writer) ClearError() { w.err = nil }
