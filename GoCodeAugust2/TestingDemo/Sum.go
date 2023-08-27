package main

import (
	"fmt"
)

func Sum(nums ...int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Sum(nums...))
}
