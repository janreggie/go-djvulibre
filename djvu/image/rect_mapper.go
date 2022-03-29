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

	rm := &RectMapper{
		src: src,
		dst: dst,
	}
	rm.precalc()
	return rm, nil
}

// Set sets the source and destination rectangles
func (rm *RectMapper) Set(src Rect, dst Rect) error {
	if src.IsEmpty() || dst.IsEmpty() {
		return errors.New("either src or dst is empty")
	}

	rm.src = src
	rm.dst = dst

	if rm.code&_SWAPXY != 0 {
		rm.src.xmin, rm.src.ymin = rm.src.ymin, rm.src.xmin
		rm.src.xmax, rm.src.ymax = rm.src.ymax, rm.src.xmax
	}

	rm.precalc()
	return nil
}

// precalc calculates rw and rh when changing/updating the src/dst rectangles
func (rm *RectMapper) precalc() {
	rm.rw = newRatio(rm.dst.Width(), rm.src.Width())
	rm.rh = newRatio(rm.dst.Height(), rm.src.Height())
}

// Get gets the source and destination rectangles
func (rm *RectMapper) Get() (src, dst Rect) {
	src = rm.src
	dst = rm.dst
	return
}

// Rotate composes the affine transform
// with a rotation of `turns` quarter turns *counter-clockwise*.
func (rm *RectMapper) Rotate(turns int) {
	oldcode := rm.code
	switch turns & 0x3 {
	case 1:
		if rm.hasFlag(_SWAPXY) {
			rm.code ^= _MIRRORY
		} else {
			rm.code ^= _MIRRORX
		}
		rm.code ^= _SWAPXY

	case 2:
		rm.code ^= (_MIRRORX | _MIRRORY)

	case 3:
		if rm.hasFlag(_SWAPXY) {
			rm.code ^= _MIRRORX
		} else {
			rm.code ^= _MIRRORY
		}
		rm.code ^= _SWAPXY
	}

	if (oldcode^rm.code)&_SWAPXY != 0 {
		rm.src.xmin, rm.src.ymin = rm.src.ymin, rm.src.xmin
		rm.src.xmax, rm.src.ymax = rm.src.ymax, rm.src.xmax
	}
}

// MirrorX composes the affine transform with a symmetry
// with respect to the vertical line
// crossing the center of the output rectangle.
func (rm *RectMapper) MirrorX() {
	rm.code ^= _MIRRORX
}

// MirrorX composes the affine transform with a symmetry
// with respect to the horizontal line
// crossing the center of the output rectangle.
func (rm *RectMapper) MirrorY() {
	rm.code ^= _MIRRORY
}

// Map maps a point according to the affine transform.
// (x,y) represents the coordinates of a point.
// This method returns the coordinates of a second point
// located in the same position relative to the corners of the output rectangle
// as the first point relative to the matching corners of the input rectangle.
// Coordinates are rounded to the nearest integer.
func (rm *RectMapper) Map(x, y int32) (int32, int32) {
	x2, y2 := x, y

	// Swap and mirror
	if rm.hasFlag(_SWAPXY) {
		x2, y2 = y2, x2
	}
	if rm.hasFlag(_MIRRORX) {
		x2 = rm.src.xmin + rm.src.xmax - x2
	}
	if rm.hasFlag(_MIRRORY) {
		y2 = rm.src.ymin + rm.src.ymax - y2
	}

	// Scale and translate
	x2 = rm.dst.xmin + rmul(x2-rm.src.xmin, rm.rw)
	y2 = rm.dst.ymin + rmul(y2-rm.src.ymin, rm.rh)

	return x2, y2
}

// MapRect returns a Rectangle
// whose points are mapped according to the affine transform.
func (rm *RectMapper) MapRect(r Rect) Rect {
	xmin, ymin := rm.Map(r.xmin, r.ymin)
	xmax, ymax := rm.Map(r.xmax, r.ymax)
	if xmin > xmax {
		xmin, xmax = xmax, xmin
	}
	if ymin > ymax {
		ymin, ymax = ymax, ymin
	}

	return Rect{
		xmin: xmin,
		ymin: ymin,
		xmax: xmax,
		ymax: ymax,
	}
}

// Unmap maps a point according to the inverse of the affine transform.
// That is, x1,y1 = rm.Unmap(x2,y2) <-> x2,y2 = rm.Map(x1,y1).
func (rm *RectMapper) Unmap(x, y int32) (int32, int32) {
	// Scale and translate
	x2 := rm.src.xmin + rdiv(x-rm.dst.xmin, rm.rw)
	y2 := rm.src.ymin + rdiv(y-rm.dst.ymin, rm.rh)

	// Mirror and swap
	if rm.hasFlag(_MIRRORX) {
		x2 = rm.src.xmin + rm.src.xmax - x2
	}
	if rm.hasFlag(_MIRRORY) {
		y2 = rm.src.ymin + rm.src.ymax - y2
	}
	if rm.hasFlag(_SWAPXY) {
		x2, y2 = y2, x2
	}

	return x2, y2
}

// UnmapRect maps a Rectangle according to the inverse of the affine transform.
// That is, r1 = rm.UnmapRect(r2) <-> r2 = rm.MapRect(r1).
func (rm *RectMapper) UnmapRect(r Rect) Rect {
	xmin, ymin := rm.Unmap(r.xmin, r.ymin)
	xmax, ymax := rm.Unmap(r.xmax, r.ymax)
	if xmin > xmax {
		xmin, xmax = xmax, xmin
	}
	if ymin > ymax {
		ymin, ymax = ymax, ymin
	}

	return Rect{
		xmin: xmin,
		ymin: ymin,
		xmax: xmax,
		ymax: ymax,
	}
}

// hasFlag checks if rm.code has some flag set
func (rm *RectMapper) hasFlag(f byte) bool {
	return rm.code&f != 0
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
