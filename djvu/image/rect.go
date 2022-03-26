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
func NewRect(xmin int32, ymin int32, width int32, height int32) Rect {
	return Rect{
		xmin: xmin,
		ymin: ymin,
		xmax: xmin + width,
		ymax: ymin + height,
	}
}

// Width returns the rectangle's width
func (r Rect) Width() int32 {
	panic("unimplemented")
}

// Height returns the rectangle's height
func (r Rect) Height() int32 {
	panic("unimplemented")
}

// Area returns the rectangle's area
func (r Rect) Area() int32 {
	panic("unimplemented")
}

// IsEmpty returns whether the rectangle is "empty"
func (r Rect) IsEmpty() bool {
	panic("unimplemented")
}

// Contains returns whether the point (x,y) is inside the rectangle.
func (r Rect) Contains(x int, y int) bool {
	panic("unimplemented")
}

// Contains whether rectangle r2 is inside the rectangle.
// That is, if the intersection of the two equals r2.
func (r Rect) ContainsRect(r2 Rect) bool {
	panic("unimplemented")
}

// Inflate increases the distance from each of the vertical sides to the centre by dx
// and each of the horizontal sides to the centre by dy,
// and returns the resulting rectangle.
// A negative value of either will push the sides closer to the centre.
// If parallel sides meet or cross each other due to negative values,
// return an empty rectangle.
func (r Rect) Inflate(dx int32, dy int32) Rect {
	panic("unimplemented")
}

// Translate moves the rectangle dx units horizontally and dy units vertically.
func (r Rect) Translate(dx int32, dy int32) Rect {
	panic("unimplemented")
}

// Intersect returns the intersection of r and r2, or an empty rectangle.
func (r Rect) Intersect(r2 Rect) Rect {
	panic("unimplemented")
}

// Recthull returns the smallest rectangle that contains all points in both rectangles.
// If either rectangle is empty, return the other one.
func (r Rect) Recthull(r2 Rect) Rect {
	panic("unimplemented")
}

// Scale expands the rectangle, relative to the origin, by a scale factor.
func (r Rect) Scale(factor float64) Rect {
	panic("unimplemented")
}

// ScaleXY expands the horizontal and vertical dimensions of the rectangle,
// relative to the origin,
// by xfactor and yfactor respectively.
func (r Rect) ScaleXY(xfactor float64, yfactor float64) Rect {
	panic("unimplemented")
}
