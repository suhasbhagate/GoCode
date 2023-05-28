package main

import(
	"fmt"
)

func main(){
	nums := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("Sum ",Sum(nums...))
	fmt.Println("Even Sum ",EvenSum(Sum,nums...))
	fmt.Println("Odd Sum ",OddSum(Sum,nums...))
}

func Sum(nums ...int)int{
	sum :=0
	for _,v:= range nums{
		sum+=v
	}
	return sum
}

func EvenSum(f func(n ...int)int, nums ...int)int{
	var enums []int
	for _,v:= range nums{
		if v%2==0{
			enums = append(enums, v)
		}
	}
	return f(enums...)
}

func OddSum(f func(n ...int)int, nums ...int)int{
	var onums []int
	for _,v:= range nums{
		if v%2!=0{
			onums = append(onums, v)
		}
	}
	return f(onums...)
}