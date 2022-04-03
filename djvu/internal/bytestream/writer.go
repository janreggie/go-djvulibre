package bytestream

import "io"

// Writer writes into a stream.
type Writer struct {
	w io.Writer
	e error
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: w,
		e: nil,
	}
}

func (w *Writer) Write8(v uint8) *Writer {
	if w.e != nil {
		return w
	}

	_, w.e = w.w.Write([]byte{v})
	return w
}

// Write16 writes a 16-bit number in big-endian format.
func (w *Writer) Write16(v uint16) *Writer {
	if w.e != nil {
		return w
	}

	bs := []byte{
		uint8((v >> 8) & 0xff),
		uint8(v & 0xff),
	}
	_, w.e = w.w.Write(bs)
	return w
}

func (w *Writer) Error() error { return w.e }

func (w *Writer) ClearError() { w.e = nil }
