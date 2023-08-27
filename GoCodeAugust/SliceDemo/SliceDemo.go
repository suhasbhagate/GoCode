package main

import(
	"fmt"
)

func main(){
	mslc := make([][]int,3)
	for i:= 0; i<len(mslc); i++{
		mslc[i] = make([]int,3)
	}
	fmt.Println(mslc)
}