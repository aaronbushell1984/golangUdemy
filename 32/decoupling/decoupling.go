// Package decoupling demonstrates decoupling using an interface
//
// The shape interface:
//	type Shape interface {
//		Area() float64
//	}
// Allows different implementations of the Area method:
//	func (c Circle) Area() float64 {
//		return math.Pi * c.radius * c.radius
//	}
//
//	func (r Rectangle) Area() float64 {
//		return r.height * r.width
//	}
//
//	func (t Triangle) Area() float64 {
//		return t.height * t.base * 0.5
//	}
// This allows parametric polymorphism by allowing Circle, Rectangle or Triangle to be passed to a Shape interface
package decoupling

import "math"

// Shape is any type which implements the Area method returning a float64
type Shape interface {
	Area() float64
}

// Rectangle defines the simplest properties of a rectangle
type Rectangle struct {
	height float64
	width  float64
}

// Circle defines the simplest properties of a circle
type Circle struct {
	radius float64
}

// Triangle defines the simplest properties of a triangle
type Triangle struct {
	height float64
	base   float64
}

// Perimeter returns the area of a rectangle
func Perimeter(r Rectangle) float64 {
	return 2 * (r.height + r.width)
}

// Area method returns the area of a Circle receiver
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Area method returns the area of a Rectangle receiver
func (r Rectangle) Area() float64 {
	return r.height * r.width
}

// Area method returns the area of a Triangle receiver
func (t Triangle) Area() float64 {
	return t.height * t.base * 0.5
}
