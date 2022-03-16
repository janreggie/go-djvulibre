package image

import (
	"sync"

	"github.com/janreggie/go-djvulibre/djvu/bytestream"
	"github.com/pkg/errors"
)

// Bitmap represents bilevel and gray-level images.
// Images are usually represented using one byte per pixel.
// Value zero represents a white pixel.
// A value equal to the number of gray levels minus one
// represents a black pixel.
// The number of gray levels is returned by the GetGrays method
// and can be manipulated by SetGrays and ChangeGrays.
//
// GetLine returns a pointer to the bytes composing one line of the image.
// This pointer can be used to read or write the image pixels.
// Line zero represents the bottom line of the image.
//
// The memory organization is setup in such a way
// that you can safely read a few pixels
// located in a small border surrounding all four sides of the image.
// The width of this border can be modified
// using the function Minborder.
// The border pixels are initialized to zero
// and therefore represent white pixels.
// You should never write anything into border pixels
// because they are shared between images and between lines.
type Bitmap struct {
	nrows       uint8
	ncols       uint8
	border      uint8
	bytesPerRow uint8
	grays       uint8
	bytes       []byte // TODO: Make this Go-like
	bytesData   []byte // TODO: Make this Go-like
	rle         []byte
	rlerows     [][]byte
	rlelength   uint16
	mtx         sync.RWMutex
	zerobuffer  *zerobuffer
}

// NewBitmap constructs an empty Bitmap object.
// It will have zero rows and columns.
func NewEmptyBitmap() *Bitmap {
	return &Bitmap{}
}

// NewBitmap constructs a Bitmap object with `nrows` rows and `ncols` columns.
// All pixels are initialized to white.
// `border` specifies the size of an optional border of white pixels
// surrounding the image.
// The number of gray levels is initially set to 2.
func NewBitmap(nrows uint8, ncols uint8, border uint8) (*Bitmap, error) {
	b := NewEmptyBitmap()
	err := b.Init(nrows, ncols, border)
	if err != nil {
		return nil, errors.Wrapf(err, "could not create an image with rows %v, cols %v, border size %v", nrows, ncols, border)
	}
	return b, nil
}

// NewBitmapFromStream creates a Bitmap
// by reading PBM, PGM or RLE data from a ByteStream.
// `border` specifies the size of an optional border of white pixels
// surrounding the image.
func NewBitmapFromStream(bs bytestream.ByteStream) (*Bitmap, error) {
	panic("unimplemented")
}

// Copy copies an existing Bitmmap and returns said copy
func (b *Bitmap) Copy() *Bitmap {
	panic("unimplemented")
}

// CopyWithBorder copies an existing Bitmmap and returns said copy,
// but sets border to a specified value.
func (b *Bitmap) CopyWithBorder(border uint8) *Bitmap {
	panic("unimplemented")
}

// CopySection creates a Bitmap by copying a rectangular segment `rect` of Bitmap `b`.
// `border` specifies the size of an optional border of white pixels
// surrounding the image.
func (b *Bitmap) CopySection(r Rect, border uint8) {
	panic("unimplemented")
}

// Init resets the Bitmap size to `nrows` by `ncols` and sets all pixels to white.
// `border` specifies the size of an optional border of white pixels
// surrounding the image.
// THe number of gray levels is initially set to 2.
func (b *Bitmap) Init(nrows uint8, ncols uint8, border uint8) error {
	// Some checking to make sure nothing overflows
	nr, nc, br := uint32(nrows), uint32(ncols), uint32(border)
	np := nr*(nc+br) + br
	if nc+br != uint32(ncols+nrows) ||
		(nr > 0 && (np-br)/nr != uint32(ncols+border)) {
		return errors.Errorf("Bitmap: image size exceeds maximum (corrupted file?)")
	}

	b.mtx.Lock()
	defer b.mtx.Unlock()
	b.destroy()
	b.grays = 2
	b.nrows = nrows
	b.ncols = ncols
	b.border = border
	b.bytesPerRow = ncols + border
	b.zerobuffer = zeroes(uint32(b.bytesPerRow) + br)

	npixels := uint32(nrows)*uint32(b.bytesPerRow) + uint32(border)
	if npixels > 0 {
		b.bytesData = make([]byte, npixels)
		b.bytes = b.bytesData
	}

	return nil
}

// Fill initializes all Bitmap pixels to some `value`
func (b *Bitmap) Fill(value uint8) {
	panic("unimplemented")
}

// Rows returns the number of rows (the image height)
func (b *Bitmap) Rows() uint32 {
	b.mtx.RLock()
	defer b.mtx.RUnlock()
	return uint32(b.nrows)
}

// Cols returns the number of columns (the image width)
func (b *Bitmap) Cols() uint32 {
	b.mtx.RLock()
	defer b.mtx.RUnlock()
	return uint32(b.ncols)
}

// destroy "resets" everything
func (b *Bitmap) destroy() {
	b.bytesData = make([]byte, 0)
	b.bytes = make([]byte, 0)
	b.rle = make([]byte, 0)
	b.rlerows = make([][]byte, 0)
	b.rlelength = 0
}
