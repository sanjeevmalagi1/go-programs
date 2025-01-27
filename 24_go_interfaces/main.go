package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.124 * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func PrintShapeDetails(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

func main() {
	var s Shape

	s = Circle{radius: 5}

	PrintShapeDetails(s)

	s = Rectangle{width: 5, height: 10}
	PrintShapeDetails(s)

}
