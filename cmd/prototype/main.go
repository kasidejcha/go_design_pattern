package main

import (
	"fmt"
)

// Shape is the prototype interface
type Shape interface {
	Clone() Shape
	GetInfo() string
}

// Circle is a concrete prototype
type Circle struct {
	Radius int
	Color  string
}

// Clone method creates a copy of the Circle
func (c *Circle) Clone() Shape {
	return &Circle{
		Radius: c.Radius,
		Color:  c.Color,
	}
}

// GetInfo returns details about the circle
func (c *Circle) GetInfo() string {
	return fmt.Sprintf("Circle: Radius=%d, Color=%s", c.Radius, c.Color)
}

// Rectangle is another concrete prototype
type Rectangle struct {
	Width  int
	Height int
	Color  string
}

// Clone method creates a copy of the Rectangle
func (r *Rectangle) Clone() Shape {
	return &Rectangle{
		Width:  r.Width,
		Height: r.Height,
		Color:  r.Color,
	}
}

// GetInfo returns details about the rectangle
func (r *Rectangle) GetInfo() string {
	return fmt.Sprintf("Rectangle: Width=%d, Height=%d, Color=%s", r.Width, r.Height, r.Color)
}

func main() {
	// Create an original circle
	circle1 := &Circle{
		Radius: 5,
		Color:  "Red",
	}
	fmt.Println(circle1.GetInfo())

	// Clone the original circle
	circle2 := circle1.Clone()
	fmt.Println(circle2.GetInfo())

	// // Modify the cloned circle's properties
	// clonedCircle := circle2.(*Circle)
	// clonedCircle.Color = "Blue"
	// fmt.Println(clonedCircle.GetInfo())
	// fmt.Println(circle1.GetInfo()) // Original circle remains unchanged

	// // Create an original rectangle
	// rectangle1 := &Rectangle{
	// 	Width:  10,
	// 	Height: 5,
	// 	Color:  "Green",
	// }
	// fmt.Println(rectangle1.GetInfo())

	// // Clone the rectangle
	// rectangle2 := rectangle1.Clone()
	// fmt.Println(rectangle2.GetInfo())
}
