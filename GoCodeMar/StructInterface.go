package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func Display(s Shape) {
	fmt.Println("Area ", s.Area())
}

func main() {
	cc := Circle{10}
	rr := Rectangle{10, 20}
	Display(cc)
	Display(rr)
}
