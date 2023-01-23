package main

import (
	"fmt"
)

func Division(a, b int){
	defer HandlePanic()
	if b==0 {
		panic("Divide by Zero")
	} else {
		fmt.Println(float64(a/b))
	}
}

func HandlePanic(){
	r := recover()
	if r!= nil{
		fmt.Println("Recovered ",r)
	}
	//fmt.Println("Recovered")
}

func main(){
	Division(6, 5)
	Division(2, 0)
}