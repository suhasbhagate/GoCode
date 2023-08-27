package main

import(
	"fmt"
)

func main(){
	var i int = 10
	var ptr1 *int = &i
	fmt.Printf("i: %v Type: %T\n", i, i)
	fmt.Printf("&i: %v Type: %T\n", &i, &i)
	fmt.Printf("ptr1: %v Type: %T\n", ptr1, ptr1)
	fmt.Printf("*ptr1: %v Type: %T\n", *ptr1, *ptr1)
	fmt.Printf("&ptr1: %v Type: %T\n", &ptr1, &ptr1)

	ptr2 := &i
	fmt.Printf("ptr2: %v Type: %T\n", ptr2, ptr2)
	fmt.Printf("*ptr2: %v Type: %T\n", *ptr2, *ptr2)
	
	*ptr1++
	fmt.Printf("*ptr1: %v Type: %T\n", *ptr1, *ptr1)
	fmt.Printf("*ptr2: %v Type: %T\n", *ptr2, *ptr2)

	*ptr2++
	fmt.Printf("*ptr1: %v Type: %T\n", *ptr1, *ptr1)
	fmt.Printf("*ptr2: %v Type: %T\n", *ptr2, *ptr2)
}