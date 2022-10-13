package interfaces

import (
	"math"
)

/**
Interfaces are named collections of method signatures.
*/

type Geometry interface {
	area() float32
	perimeter() float32
}

type Rectangle struct {
	width, height float32
}

type Circle struct {
	radius float32
}

func (r Rectangle) area() float32 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float32 {
	return 2*r.width + 2*r.height
}

func (c Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func measureArea(g Geometry) float32 {
	return g.area()
}

func measurePerimeter(g Geometry) float32 {
	return g.perimeter()
}
