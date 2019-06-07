/*
Package creational provides various object creation mechanisms, which increase flexibility and reuse of existing code.
*/
package factory

import (
	"fmt"
	"io"
	"os"
)

type Document struct {
	ShapeFactories []ShapeFactory
}

func (doc *Document) Draw(w io.Writer) error {
	viewport := Viewport{
		Location: Point{
			X: 0,
			Y: 0,
		},
		Size: Size{
			Width:  640,
			Height: 480,
		},
	}
	if _, err := fmt.Fprintf(w, `<svg height="%f" width="%f">`, viewport.Size.Height, viewport.Size.Width); err != nil {
		return err
	}
	for _, factory := range doc.ShapeFactories {
		shape := factory.Create(viewport)
		if err := shape.Draw(w); err != nil {
			return err
		}
	}
	_, err := fmt.Fprint(w, `</svg>`)
	return err
}

func Init() {
	fmt.Println("Executing factory pattern...\n")
	doc := &Document{
		ShapeFactories: []ShapeFactory{
			&CircleFactory{},
			&RectangleFactory{},
		},
	}
	doc.Draw(os.Stdout)
	fmt.Println("\n\nFactory pattern completed!")
}
