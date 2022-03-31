package image

// Rect represents a rectangle
// whose sides are parallel to the axes.
// Such a rectangle represents all the points
// whose coordinates lies between well defined minimal and maximal values.
// Methods can combine several rectangles
// by computing the intersection of rectangles (Intersect)
// or the smallest rectangle enclosing two rectangles (Recthull).
//
// A rectangle is represented by two "points" (xmin,ymin) and (xmax,ymax).
// A rectangle contains all pixels
// with horizontal pixel coordinates in range xmin (inclusive) to xmax (inclusive)
// and vertical coordinates ymin (inclusive) to ymax (exclusive).
// That is, (xmax,ymax) is *not* inside the rectangle.
//
// The zero value returns a rectangle
// with height and width 0
// and whose corners lie on (0,0)
// i.e., the bottom-left "corner".
type Rect struct {
	// Minimal horizontal point coordinate of the rectangle.
	xmin int32
	// Minimal vertical point coordinate of the rectangle.
	ymin int32
	// Maximal horizontal point coordinate of the rectangle.
	xmax int32
	// Maximal vertical point coordinate of the rectangle.
	ymax int32
}

// NewRect returns a Rect with minimal coordinates (xmin,ymin)
// with respective width and height.
// Width and height can be negative,
// but the internal structure will adjust itself.
func NewRect(xmin int32, ymin int32, width int32, height int32) Rect {
	return Rect{
		xmin: min(xmin, xmin+width),
		ymin: min(ymin, ymin+height),
		xmax: max(xmin, xmin+width),
		ymax: max(ymin, ymin+height),
	}
}

// Width returns the rectangle's width
func (r Rect) Width() int32 {
	return r.xmax - r.xmin
}

// Height returns the rectangle's height
func (r Rect) Height() int32 {
	return r.ymax - r.ymin
}

// Area returns the rectangle's area
func (r Rect) Area() int32 {
	if r.IsEmpty() {
		return 0
	}
	return r.Width() * r.Height()
}

// IsEmpty returns whether the rectangle is "empty"
func (r Rect) IsEmpty() bool {
	return isEmpty(r.xmin, r.xmax, r.ymin, r.ymax)
}

func isEmpty(xmin, xmax, ymin, ymax int32) bool {
	return xmin >= xmax || ymin >= ymax
}

// Contains returns whether the point (x,y) is inside the rectangle.
func (r Rect) Contains(x int32, y int32) bool {
	return x >= r.xmin && x < r.xmax &&
		y >= r.ymin && y < r.ymax
}

// Contains whether rectangle r2 is inside the rectangle.
// That is, if the intersection of the two equals r2.
func (r Rect) ContainsRect(r2 Rect) bool {
	return r.Intersect(r2) == r2
}

// Inflate increases the distance from each of the vertical sides to the centre by dx
// and each of the horizontal sides to the centre by dy,
// and returns the resulting rectangle.
// A negative value of either will push the sides closer to the centre.
// If parallel sides meet or cross each other due to negative values,
// return an empty rectangle.
func (r Rect) Inflate(dx int32, dy int32) Rect {
	xmin, xmax := r.xmin-dx, r.xmax+dx
	ymin, ymax := r.ymin-dy, r.ymax+dy
	return sanitize(xmin, xmax, ymin, ymax)
}

// Translate moves the rectangle dx units horizontally and dy units vertically.
func (r Rect) Translate(dx int32, dy int32) Rect {
	xmin, xmax := r.xmin+dx, r.xmax+dx
	ymin, ymax := r.ymin+dy, r.ymax+dy
	return sanitize(xmin, xmax, ymin, ymax)
}

// Intersect returns the intersection of r and r2, or an empty rectangle.
func (r Rect) Intersect(r2 Rect) Rect {
	xmin := max(r.xmin, r2.xmin)
	xmax := min(r.xmax, r2.xmax)
	ymin := max(r.ymin, r2.ymin)
	ymax := min(r.ymax, r2.ymax)
	return sanitize(xmin, xmax, ymin, ymax)
}

// Recthull returns the smallest rectangle that contains all points in both rectangles.
// If either rectangle is empty, return the other one.
func (r Rect) Recthull(r2 Rect) Rect {
	panic("unimplemented")
}

// Scale expands the rectangle,
// relative to the origin,
// by a scale factor.
// Putting a negative factor
// will return an empty rectangle.
func (r Rect) Scale(factor float64) Rect {
	return r.ScaleXY(factor, factor)
}

// ScaleXY expands the horizontal and vertical dimensions of the rectangle,
// relative to the origin,
// by xfactor and yfactor respectively.
// Putting a negative xfactor or yfactor
// will return an empty rectangle.
func (r Rect) ScaleXY(xfactor float64, yfactor float64) Rect {
	if xfactor <= 0 || yfactor <= 0 {
		return Rect{}
	}

	return Rect{
		xmin: int32(float64(r.xmin) * xfactor),
		xmax: int32(float64(r.xmax) * xfactor),
		ymin: int32(float64(r.ymin) * yfactor),
		ymax: int32(float64(r.ymax) * yfactor),
	}
}

// sanitize returns an empty rectangle
// if the input bounds are correct,
// and the appropriate rectangle otherwise.
func sanitize(xmin, xmax, ymin, ymax int32) Rect {
	if isEmpty(xmin, xmax, ymin, ymax) {
		return Rect{}
	}
	return Rect{
		xmin: xmin,
		ymin: ymin,
		xmax: xmax,
		ymax: ymax,
	}
}

func min(x, y int32) int32 {
	if x > y {
		return y
	}
	return x
}

func max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}
