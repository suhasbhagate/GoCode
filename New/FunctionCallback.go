package main

import (
	"fmt"
)

func main(){
	arr := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(sum(arr...))
	fmt.Println(EvenSum(sum,arr...))
	fmt.Println(OddSum(sum, arr...))

}

func sum(nums ...int)int{
	s := 0
	for _, v := range nums{
		s += v
	}
	return s
}

func EvenSum(f func(enums ...int)int, nums ...int)int {
	var enums []int
	for _,v := range nums{
		if v % 2 ==0{
			enums = append(enums, v)
		}
	}
	return f(enums...)
}

func OddSum(f func(onums ...int)int, nums ...int)int {
	var onums []int
	for _,v := range nums{
		if v % 2 !=0{
			onums = append(onums, v)
		}
	}
	return f(onums...)
}