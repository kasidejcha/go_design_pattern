package main

import "fmt"

// Implementer: Renderer interface defines drawing methods
type Renderer interface {
	RenderCircle(radius float64)
}

// Concrete Implementer 1: RasterRenderer
type RasterRenderer struct{}

func (r *RasterRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing a raster circle with radius %f\n", radius)
}

// Concrete Implementer 2: VectorRenderer
type VectorRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing a vector circle with radius %f\n", radius)
}


// Abstraction: Shape
type Shape struct {
	renderer Renderer
}

func (s *Shape) Draw() {
	// Placeholder method, should be overridden by refined abstractions
}

func (s *Shape) SetRenderer(renderer Renderer) {
	s.renderer = renderer
}

// Refined Abstraction: Circle
type Circle struct {
	Shape  // Embedding the Shape (abstraction)
	radius float64
}

func NewCircle(renderer Renderer, radius float64) *Circle {
	return &Circle{
		Shape:  Shape{renderer: renderer},
		radius: radius,
	}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func main() {
	// Using RasterRenderer
	rasterRenderer := &RasterRenderer{}
	circle := NewCircle(rasterRenderer, 5.0)
	circle.Draw()

	// Switch to VectorRenderer
	vectorRenderer := &VectorRenderer{}
	circle.SetRenderer(vectorRenderer)
	circle.Draw()
}
