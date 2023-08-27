package main

import(
	"fmt"
)

func main(){
	slc := make([][]int, 3)
	for i := range slc{
		slc[i] = make([]int, 3)
	}
	fmt.Println(slc)
}