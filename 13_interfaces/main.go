package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	w, h float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (x Rectangle) area() float64 {
	return x.w * x.h
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, r: 5}
	rectangle := Rectangle{w: 10, h: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}
