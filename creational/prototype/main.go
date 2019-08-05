package prototype

import (
	"fmt"
	"reflect"
)

func Init() {
	var shapes []Cloner
	var cloneShapes []Cloner

	circle := Circle{
		Shape:  Shape{10, 20, "red"},
		radius: 15,
	}

	shapes = make([]Cloner, 0)
	shapes = append(shapes, &circle)

	anotherCircle := circle.clone()
	shapes = append(shapes, anotherCircle)

	rectangle := Rectangle{
		Shape: Shape{10, 20, ""},
	}
	shapes = append(shapes, &rectangle)

	cloneShapes = make([]Cloner, 0)
	cloneAndCompare(shapes, cloneShapes)
}

func cloneAndCompare(shapes []Cloner, cloneShapes []Cloner) {
	for _, shape := range shapes {
		cloneShapes = append(cloneShapes, shape.clone())
	}

	for idx := 0; idx < len(shapes); idx++ {
		if shapes[idx] != cloneShapes[idx] {
			fmt.Printf("%d: Shapes are different objects (yay!)\n", idx)
			if reflect.DeepEqual(shapes[idx], cloneShapes[idx]) {
				fmt.Printf("%d: And they are identical (yay!)\n", idx)
			} else {
				fmt.Printf("%d: But they are not identical (boo!)\n", idx)
			}
		} else {
			fmt.Printf("%d: Shape objects are the same (boo!)\n", idx)
		}
	}
}
