package image

import (
	"io"
	"math"

	"github.com/janreggie/go-djvulibre/djvu/internal/bytestream"
	"github.com/pkg/errors"
)

// NewBitmapFromStream creates a Bitmap
// by reading PBM, PGM or RLE data from a ByteStream.
// `border` specifies the size of an optional border of white pixels
// surrounding the image.
func NewBitmapFromStream(r io.Reader, border uint16) (*Bitmap, error) {
	magic := make([]byte, 2)
	if _, err := io.ReadAtLeast(r, magic, 2); err != nil {
		return nil, errors.Wrapf(err, "could not read magic number")
	}
	br := bytestream.NewReader(r)

	acols, err := br.ReadInteger()
	if err != nil {
		return nil, errors.Wrapf(err, "could not read number of columns")
	}
	if acols > math.MaxUint16 {
		return nil, errors.Wrapf(err, "too many columns in image")
	}
	arows, err := br.ReadInteger()
	if err != nil {
		return nil, errors.Wrapf(err, "could not read number of rows")
	}
	if arows > math.MaxUint16 {
		return nil, errors.Wrapf(err, "too many rows in image")
	}
	bitmap, err := NewBitmap(uint16(arows), uint16(acols), border)
	if err != nil {
		return nil, errors.Wrapf(err, "could not init bitmap")
	}

	switch magic[0] {
	case 'P':
		switch magic[1] {
		case '1':
			bitmap.grays = 2
			if err := bitmap.readPbmText(r); err != nil {
				return nil, errors.Wrapf(err, "could not read PBM text")
			}
			return bitmap, nil

		case '2':
			maxval, err := br.ReadInteger()
			if err != nil {
				return nil, errors.Wrapf(err, "could not read maxval for PGM text")
			}
			if maxval > 65535 {
				return nil, errors.Errorf("cannot read PGM with depth greater than 16 bits, got depth of %v", maxval)
			}
			if maxval > 255 {
				bitmap.grays = 256
			} else {
				bitmap.grays = uint16(maxval) + 1
			}
			if err := bitmap.readPgmText(r, maxval); err != nil {
				return nil, errors.Wrapf(err, "could not read PGM text")
			}
			return bitmap, nil

		case '4':
			bitmap.grays = 2
			if err := bitmap.readPbmRaw(r); err != nil {
				return nil, errors.Wrapf(err, "could not read PBM raw")
			}
			return bitmap, nil

		case '5':
			maxval, err := br.ReadInteger()
			if err != nil {
				return nil, errors.Wrapf(err, "could not read maxval for PGM raw")
			}
			if maxval > 65535 {
				return nil, errors.Errorf("cannot read PGM with depth greater than 16 bits, got depth of %v", maxval)
			}
			if maxval > 255 {
				bitmap.grays = 256
			} else {
				bitmap.grays = uint16(maxval) + 1
			}
			if err := bitmap.readPgmRaw(r, maxval); err != nil {
				return nil, errors.Wrapf(err, "could not read PGM raw")
			}
			return bitmap, nil
		}

	case 'R':
		switch magic[1] {
		case '4':
			bitmap.grays = 2
			if err := bitmap.readRleRaw(r); err != nil {
				return nil, errors.Wrapf(err, "could not read RLE raw")
			}
			return bitmap, nil
		}
	}

	return nil, errors.Errorf("invalid magic bytes %v", magic)
}

func (b *Bitmap) readPbmText(r io.Reader) error {
	br := bytestream.NewByteReader(r)

	for rr := b.nrows - 1; rr <= b.nrows-1; rr-- {
		row := rr*b.bytesPerRow + b.border

		for cc := uint16(0); cc < b.ncols; cc++ {
			if br.Advance(); br.E != nil {
				return br.E
			}
			for br.B == ' ' || br.B == '\t' || br.B == '\r' || br.B == '\n' {
				if br.Advance(); br.E != nil {
					return br.E
				}
			}

			switch br.B {
			case '1':
				b.bytes[row+cc] = 1
			case '0':
				b.bytes[row+cc] = 0
			default:
				return errors.Errorf("bad PBM")
			}
		}
	}

	return nil
}

func (b *Bitmap) readPgmText(r io.Reader, maxval uint32) error {
	ramp := make([]byte, maxval+1)
	for ii := uint32(0); ii < maxval; ii++ {
		ramp[ii] = byte((uint32(b.grays-1)*(maxval-ii) + maxval/2) / maxval)
	}

	br := bytestream.NewReader(r)

	for rr := b.nrows - 1; rr <= b.nrows-1; rr-- {
		row := rr*b.bytesPerRow + b.border

		for cc := uint16(0); cc < b.ncols; cc++ {
			ind, err := br.ReadInteger()
			if err != nil {
				return err
			}
			b.bytes[row+cc] = ramp[ind]
		}
	}

	return nil
}

func (b *Bitmap) readPbmRaw(r io.Reader) error {
	br := newBytereader(r)

	for rr := b.nrows - 1; rr <= b.nrows-1; rr-- {
		row := rr*b.bytesPerRow + b.border
		var mask byte
		for cc := uint16(0); cc < b.ncols; cc++ {
			if mask == 0 {
				if br.advance(); br.E != nil {
					return br.E
				}
				mask = 0x80
			}
			if br.B&mask != 0 {
				b.bytes[row+cc] = 1
			} else {
				b.bytes[row+cc] = 0
			}
			mask >>= 1
		}
	}

	return nil
}

func (b *Bitmap) readPgmRaw(r io.Reader, maxval uint32) error {
	br := newBytereader(r)
	maxbin := uint32(256)
	if maxval > 255 {
		maxbin = 65536
	}

	ramp := make([]byte, maxbin-1)
	for ii := uint32(0); ii < maxval; ii++ {
		ramp[ii] = byte((uint32(b.grays-1)*(maxval-ii) + maxval/2) / maxval)
	}

	for rr := b.nrows - 1; rr <= b.nrows-1; rr-- {
		row := rr*b.bytesPerRow + b.border
		for cc := uint16(0); cc < b.ncols; cc++ {
			if maxbin > 255 { // Two bytes read
				if br.advance(); br.E != nil {
					return br.E
				}
				b1 := br.B
				if br.advance(); br.E != nil {
					return br.E
				}
				b2 := br.B

				b.bytes[row+cc] = ramp[uint16(b1)*256+uint16(b2)]

			} else { // One byte read
				if br.advance(); br.E != nil {
					return br.E
				}
				bb := br.B
				b.bytes[row+cc] = ramp[bb]
			}
		}
	}

	return nil
}

func (b *Bitmap) readRleRaw(r io.Reader) error {
	br := newBytereader(r)
	var p byte
	n := b.nrows - 1
	row := n*b.bytesPerRow + b.border
	var c uint32

	for n <= b.nrows-1 {
		if br.advance(); br.E != nil {
			return br.E
		}
		x := uint32(br.B)
		if x >= _RUNOVERFLOWVALUE {
			if br.advance(); br.E != nil {
				return br.E
			}
			x2 := br.B
			x = uint32(x2) + ((x - _RUNOVERFLOWVALUE) << 8)
		}

		if c+x > uint32(b.ncols) {
			return errors.New("bitmap lost sync")
		}

		for ; x > 0; x-- {
			b.bytes[uint32(row)+c] = p
			c++
		}
		p = 1 - p
		if c >= uint32(b.ncols) {
			c = 0
			p = 0
			row -= b.bytesPerRow
			n--
		}
	}

	return nil
}
