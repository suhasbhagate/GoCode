package main

import(
	"fmt"
)


func main(){
	fmt.Println(Factorial(5))
	fmt.Println(Factorial(0))
	fmt.Println(Factorial(-5))
}


func Factorial(n int)int{
	if n<0{
		return -1
	}
	if n==0{
		return 1
	} else{
		return n * Factorial(n-1)
	}
}