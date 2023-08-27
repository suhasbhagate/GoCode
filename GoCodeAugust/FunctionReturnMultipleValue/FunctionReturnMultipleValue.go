package main

import(
	"fmt"
)

func swap(i, j int)(int,int){
	return j,i
}

func main(){
	a := 10
	b := 20
	_, _ = swap(a, b)
	fmt.Printf("a: %v\n",a)
	fmt.Printf("b: %v\n",b)
}