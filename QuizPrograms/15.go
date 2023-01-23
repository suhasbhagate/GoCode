package main

import(
	"fmt"
)

func main(){
	var x []int // a nil slice
	x = append(x, 10, 20, 30) // Verify append work or not
	fmt.Println(x)
}