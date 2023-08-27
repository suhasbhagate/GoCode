package main

import (
	"fmt"
)

func HandlePanic() {
	r := recover()
	if r != nil {
		fmt.Println("r: \n", r)
	}
	fmt.Println("Recovered")
}

func Division(num1, num2 int) {
	if num2 == 0 {
		defer HandlePanic()
		panic("Can't Divide by Zero")
	} else {
		fmt.Println(num1 / num2)
	}
}

func main() {
	Division(10, 5)
	Division(10, 4)
	Division(10, 0)
}
