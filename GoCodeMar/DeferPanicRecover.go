package main

import(
	"fmt"
)

func main(){
	fmt.Println(Divide(10,2))
	fmt.Println(Divide(12,3))
	fmt.Println(Divide(5,0))
	fmt.Println(Divide(15,5))
}

func HandlePanic(){
	a := recover()
	if a!=nil{
		fmt.Println(a)
		fmt.Println("Recovered")
	}	
}

func Divide(a, b int)int{
	defer HandlePanic()
	if b!=0{
		return a/b
	} else{
		panic("Divide by Zero")
	}
}

