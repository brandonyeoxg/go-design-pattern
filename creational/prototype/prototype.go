package prototype

type Cloner interface {
	clone() Cloner
}

type Shape struct {
	x     int
	y     int
	color string
}

type Rectangle struct {
	Shape
	width  int
	height int
}

func (rect *Rectangle) clone() Cloner {
	return &Rectangle{Shape: rect.Shape, width: rect.width, height: rect.height}
}

type Circle struct {
	Shape
	radius int
}

func (circle *Circle) clone() Cloner {
	return &Circle{Shape: circle.Shape, radius: circle.radius}
}
