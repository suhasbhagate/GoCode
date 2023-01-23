package main

import (
	"fmt"
)

func main(){
	a := Incrementer()
	b := Incrementer()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func Incrementer() func()int{
	var x int
	return func()int{
		x++
		return x
	}
}