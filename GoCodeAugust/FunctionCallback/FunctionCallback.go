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

func OddSum(f func(oddnums ...int) int, nums ...int) int {
	var oddnums []int
	for _, v := range nums {
		if v%2 == 0 {
			oddnums = append(oddnums, v)
		}
	}
	return f(oddnums...)
}

func EvenSum(f func(evennums ...int) int, nums ...int) int {
	var evennums []int
	for _, v := range nums {
		if v%2 != 0 {
			evennums = append(evennums, v)
		}
	}
	return f(evennums...)
}

func main() {
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Sum(slc...))
	fmt.Println(OddSum(Sum,slc...))
	fmt.Println(EvenSum(Sum,slc...))
}
