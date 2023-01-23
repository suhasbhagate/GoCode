package main

import(
	"fmt"
)

func main(){
	m := []int{1,2,3,4,5,6,7,8,9}
	for i, v := range m{
		fmt.Printf("Index: %v Value: %v\n", i, v)
	}
}