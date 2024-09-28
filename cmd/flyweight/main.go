package main

import "fmt"

// Flyweight: Circle with intrinsic state (color)
type Circle struct {
	color string // Intrinsic state
}

func (c *Circle) Draw(x, y, radius int) {
	fmt.Printf("Drawing circle of color %s at (%d, %d) with radius %d\n", c.color, x, y, radius)
}

// Flyweight Factory: CircleFactory
type CircleFactory struct {
	circleMap map[string]*Circle
}

func NewCircleFactory() *CircleFactory {
	return &CircleFactory{
		circleMap: make(map[string]*Circle),
	}
}

func (f *CircleFactory) GetCircle(color string) *Circle {
	if circle, exists := f.circleMap[color]; exists {
		return circle
	}

	// Create a new circle and add it to the map
	circle := &Circle{color: color}
	f.circleMap[color] = circle
	fmt.Printf("Created new circle of color: %s\n", color)
	return circle
}

func main() {
	// Create a new circle factory
	factory := NewCircleFactory()

	// Use the factory to get circles and draw them
	redCircle1 := factory.GetCircle("red")
	redCircle1.Draw(10, 10, 5)

	redCircle2 := factory.GetCircle("red")
	redCircle2.Draw(20, 20, 10)

	blueCircle := factory.GetCircle("blue")
	blueCircle.Draw(30, 30, 15)

	// Attempt to get another "red" circle
	redCircle3 := factory.GetCircle("red")
	redCircle3.Draw(40, 40, 20)
}
