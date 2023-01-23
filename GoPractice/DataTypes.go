package main

import (
	"fmt"
)

func main(){
	a:= 10
	fmt.Printf("Value %v Type %T\n",a,a)

	//a = "Suhas"
	a = 20
	fmt.Printf("Value %v Type %T\n",a,a)

	//Bit Opeartors

	x := 1
	y := 4

	fmt.Printf("AND of %b and %b is: %b\n",x, y, x & y)
	fmt.Printf("OR of %b and %b is: %b\n",x, y, x | y)
	fmt.Printf("XOR of %b and %b is: %b\n",x, y, x ^ y)
	fmt.Printf("Bit Clear of %b and %b is: %b\n",x, y, x &^ y)
	
	fmt.Printf("Left Shift of %b is: %b\n",x, x<<1)
	fmt.Printf("Right Shift of %b is: %b\n",x, x>>1)
	fmt.Printf("Not of %b is: %b",x, ^x)

}