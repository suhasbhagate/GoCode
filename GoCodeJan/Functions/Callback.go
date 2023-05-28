package main

import(
	"fmt"
)

func main(){
	slc := []int{1,2,3,4,5,6,7,8,9}
	sm := Sum(slc...)
	fmt.Println("Sum ",sm)
	osm := OddSum(Sum,slc...)
	fmt.Println("Odd Sum ",osm)
	esm := EvenSum(Sum,slc...)
	fmt.Println("Even Sum ",esm)
}

func Sum(nums ...int)int{
	sm := 0
	for _, v := range nums{
		sm += v
	}
	return sm
}

func OddSum(f func(onums ...int)int, nums ...int)int{
	var onums []int
	for _, v:= range nums{
		if v%2 != 0{
			onums = append(onums, v)
		}
	}
	return f(onums...)
}

func EvenSum(f func(enums ...int)int, nums ...int)int{
	var enums []int
	for _, v:= range nums{
		if v%2 ==0{
			enums = append(enums, v)
		}
	}
	return f(enums...)
}