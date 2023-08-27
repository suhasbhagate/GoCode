package main

import(
	"fmt"
)

func HandlePanic(){
	a := recover()
	if a!= nil{
		fmt.Println(a)
	}
	fmt.Println("Recovered")
}

func Division(i, j int){
	defer HandlePanic()
	if j==0{
		panic("Divide By Zero")
	} else{
		fmt.Println(i/j)
	}
}

func main(){
	Division(12, 4)
	Division(12, 0)
	Division(8, 4)
}