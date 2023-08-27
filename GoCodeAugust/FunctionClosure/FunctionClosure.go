package main

import (
	"fmt"
)

func Incrementer() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	a := Incrementer()
	b := Incrementer()

	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	val := a()
	fmt.Println(val)
}
