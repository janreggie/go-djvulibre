package iw44

import "github.com/janreggie/go-djvulibre/djvu/image"

// Map represents all blocks of an image
type Map struct {
	iw, ih uint16
	bw, bh uint16
	nb     uint16
	blocks *image.Block
	top    uint16
}

func NewMap(w, h uint16) *Map {
	bw := (w + 0x20 - 1) & (0 ^ 0x1f)
	bh := (h + 0x20 - 1) & (0 ^ 0x1f)
	return &Map{
		iw:     w,
		ih:     h,
		bw:     bw,
		bh:     bh,
		nb:     (bw * bh) / (32 * 32),
		blocks: &image.Block{}, // TODO: Construct
		top:    IWALLOCSIZE,
	}
}
