package main

import "fmt" 

type Number int

func (n Number) Positive() bool {
	return n > 0
}

func main() {
	num := Number(50)
	fmt.Printf("%d is positive: %t\n", num, num.Positive())
}