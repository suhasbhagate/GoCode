package main

import (
	"fmt"
	"math"
)

type Shape interface{
	Area() float64
}

type Square struct{
	side float64
}

func (s Square) Area() float64{
	return s.side * s.side
}

type Circle struct{
	radius float64
}

func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
}

func Display(s Shape){
	fmt.Println(s.Area())
}

func main(){
	s := Square{4}
	c := Circle{5}
	Display(s)
	Display(c)
}
