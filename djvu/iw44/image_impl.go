package iw44

import (
	"io"

	"github.com/janreggie/go-djvulibre/djvu/iff"
	"github.com/janreggie/go-djvulibre/djvu/image"
)

// image_impl is a concrete implementation of the Image interface
type image_impl struct {
	dbFrac  float64
	ymap    *Map
	cbmap   *Map
	crmap   *Map
	cslice  uint16
	cserial uint16
	cbytes  uint16
}

// GetWidth implements Image
func (image_impl) GetWidth() uint16 {
	panic("unimplemented")
}

// GetHeight implements Image
func (image_impl) GetHeight() uint16 {
	panic("unimplemented")
}

// GetBitmap implements Image
func (image_impl) GetBitmap() *image.Bitmap {
	panic("unimplemented")
}

// GetBitmapSample implements Image
func (image_impl) GetBitmapSample(subsample uint8, rect *image.Rect) *image.Bitmap {
	panic("unimplemented")
}

// GetPixmap implements Image
func (image_impl) GetPixmap() *image.Pixmap {
	panic("unimplemented")
}

// GetPixmapSample implements Image
func (image_impl) GetPixmapSample(subsample uint8, rect *image.Rect) *image.Pixmap {
	panic("unimplemented")
}

// GetMemoryUsage implements Image
func (image_impl) GetMemoryUsage() uintptr {
	panic("unimplemented")
}

// GetPercentMemory implements Image
func (image_impl) GetPercentMemory() uint64 {
	panic("unimplemented")
}

// EncodeChunk implements Image
func (image_impl) EncodeChunk(w io.Writer) int {
	panic("unimplemented")
}

// EncodeIFF implements Image
func (image_impl) EncodeIFF(iff *iff.IFF, chunks uint16, params *EncoderParams) {
	panic("unimplemented")
}

// DecodeChunk implements Image
func (image_impl) DecodeChunk(r io.Reader) int {
	panic("unimplemented")
}

// DecodeIff implements Image
func (image_impl) DecodeIff(iff *iff.IFF, maxChunks uint16) {
	panic("unimplemented")
}

// CloseCodec implements Image
func (image_impl) CloseCodec() {
	panic("unimplemented")
}

// GetSerial implements Image
func (image_impl) GetSerial() uint16 {
	panic("unimplemented")
}

// SetCrcbDelay implements Image
func (image_impl) SetCrcbDelay(delay uint16) uint16 {
	panic("unimplemented")
}

// SetDbFrac implements Image
func (image_impl) SetDbFrac(frac float64) {
	panic("unimplemented")
}

var _ Image = image_impl{}
