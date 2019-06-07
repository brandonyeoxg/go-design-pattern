package factory

import "math"

type ShapeFactory interface {
	Create(viewport Viewport) Shape
}

type CircleFactory struct{}

func (factory *CircleFactory) Create(viewport Viewport) Shape {
	return &Circle{
		Location: viewport.Location,
		Radius:   math.Min(viewport.Size.Width, viewport.Size.Height),
	}
}

type RectangleFactory struct{}

func (factory *RectangleFactory) Create(viewport Viewport) Shape {
	return &Rectangle{
		Location: viewport.Location,
		Size:     viewport.Size,
	}
}
