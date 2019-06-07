package factory

import (
	"fmt"
	"io"
)

type Shape interface {
	Draw(io.Writer) error
}

type Point struct {
	X float64
	Y float64
}

type Size struct {
	Width  float64
	Height float64
}

type Viewport struct {
	Location Point
	Size     Size
}

type Circle struct {
	Location Point
	Radius   float64
}

func (c *Circle) Draw(w io.Writer) error {
	_, err := fmt.Fprintf(w, `<circle cx="%f" cy="%f" r="%f"/>`, c.Location.X, c.Location.Y, c.Radius)
	return err
}

type Rectangle struct {
	Location Point
	Size     Size
}

func (rect *Rectangle) Draw(w io.Writer) error {
	_, err := fmt.Fprintf(w, `<rect x="%f" y="%f" width="%f" height="%f"/>`, rect.Location.X, rect.Location.Y, rect.Size.Width, rect.Size.Height)
	return err
}
