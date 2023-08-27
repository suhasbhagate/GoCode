package main

import (
	"fmt"
)

func main() {
	a := 10
	fmt.Printf("a: Type: %T Value: %v\n", a, a)
	fmt.Printf("&a: Type: %T Value: %v\n", &a, &a)

	ptr := &a
	fmt.Printf("ptr: Type: %T Value: %v\n", ptr, ptr)
	fmt.Printf("&ptr: Type: %T Value: %v\n", &ptr, &ptr)
	fmt.Printf("*ptr: Type: %T Value: %v\n", *ptr, *ptr)

	ptr1 := new(int)
	fmt.Printf("*ptr1: Type: %T Value: %v\n", *ptr1, *ptr1)

	ptr2 := ptr1
	fmt.Printf("*ptr2: Type: %T Value: %v\n", *ptr2, *ptr2)

	*ptr1 = 5
	fmt.Printf("*ptr1: Type: %T Value: %v\n", *ptr1, *ptr1)
	fmt.Printf("*ptr2: Type: %T Value: %v\n", *ptr2, *ptr2)

	*ptr2 = 15
	fmt.Printf("*ptr1: Type: %T Value: %v\n", *ptr1, *ptr1)
	fmt.Printf("*ptr2: Type: %T Value: %v\n", *ptr2, *ptr2)
}
