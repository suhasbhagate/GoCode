package main

import(
	"fmt"
)

func main(){
	sum:=func(i,j int)int{
		return i+j
	}(5,10)
	fmt.Println(sum)

	vr:= func(i,j int){
		fmt.Println("Sum: ",i+j)
	}
	vr(7,8)
}