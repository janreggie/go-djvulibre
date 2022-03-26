package image

import "github.com/pkg/errors"

// RectMapper maps points from one rectangle to another.
// Given the coordinates of a point in the first rectangle,
// RectMapper computes the corresponding point in the second rectangle.
//
// This function actually implements an affine transform
// which maps the corners of the first rectangle
// onto the matching corners of the second.
type RectMapper struct {
	src Rect
	dst Rect

	// For determining the type of operations from src to dst
	code byte

	// Ratios between the rectangles' width and heights
	rw ratio
	rh ratio
}

// NewRectMapper returns a new RectMapper with input rectangle src and output rectangle dst.
// This returns an error if either rectangle is empty.
func NewRectMapper(src Rect, dst Rect) (*RectMapper, error) {
	if src.IsEmpty() || dst.IsEmpty() {
		return nil, errors.New("either src or dst is empty")
	}

	return &RectMapper{
		src:  src,
		dst:  dst,
		code: 0,
		rw:   newRatio(dst.Width(), src.Width()),
		rh:   newRatio(dst.Height(), src.Height()),
	}, nil
}

// Some codes for RectMapper
const (
	_MIRRORX = 1 << iota
	_MIRRORY
	_SWAPXY
)

// ratio struct for precise computation.
type ratio struct{ p, q int32 }

func newRatio(p, q int32) ratio {
	if q == 0 {
		// Note: This should never happen within the program
		panic("division by zero in ratio")
	}

	if p == 0 {
		q = 1
	}

	if q < 0 {
		p, q = -p, -q
	}

	// Simplify the ratio
	gcd := int32(1)
	g1 := p
	g2 := q
	if g1 > g2 {
		g1, g2 = g2, g1
	}
	for g1 > 0 {
		gcd = g1
		g1 = g2 % g1
		g2 = gcd
	}
	p /= gcd
	q /= gcd

	return ratio{p: p, q: q}
}

// rmul(n,r) == n*r
func rmul(n int32, r ratio) int32 {
	x := int64(n) * int64(r.p)
	halfq := int64(r.q / 2)
	if x >= 0 {
		return int32((halfq + x) / int64(r.q))
	}

	return -int32((halfq - x) / int64(r.q))
}

// rdiv(n,r) == n/r
func rdiv(n int32, r ratio) int32 {
	x := int64(n) * int64(r.q)
	halfp := int64(r.p / 2)
	if x >= 0 {
		return int32((halfp + x) / int64(r.p))
	}

	return -int32((halfp - x) / int64(r.p))
}
