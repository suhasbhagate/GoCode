package main

import (
	"fmt"
	"math"
)

type Shape interface{
	Area() float64
}

type Circle struct{
	radius float64
}

func (c Circle) Area()float64{
	return math.Pi * c.radius * c.radius
}

type Square struct{
	side float64
}

func (s Square) Area()float64{
	return s.side * s.side
}

func Display(sh Shape){
	fmt.Printf("Shape %T Area %v\n",sh, sh.Area())
}

func main(){
	c := Circle{5}
	s := Square{6}
	Display(c)
	Display(s)
}