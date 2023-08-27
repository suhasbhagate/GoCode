package main

import(
	"fmt"
)

func Incrementer() func()int{
	var x int
	return func()int{
		x++
		return x
	}
}

func main(){
	a := Incrementer()
	b := Incrementer()
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())

	retuval := a()
	fmt.Println(retuval)

	nums := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(Sum(nums...))
	fmt.Println(EvenSum(Sum, nums...))
	fmt.Println(OddSum(Sum,nums...))
}

func Sum(nums ...int)int{
	s := 0
	for _,v := range nums{	
		s +=v
	}
	return s
}

func EvenSum(f func(evennums ...int)int, nums ...int)int{
	var evennums []int
	for _,v := range nums{
		if v%2 == 0{
			evennums = append(evennums, v)
		}
	}
	return f(evennums...)
}


func OddSum(f func(oddnums ...int)int, nums ...int)int{
	var oddnums []int
	for _,v := range nums{
		if v%2 != 0{
			oddnums = append(oddnums, v)
		}
	}
	return f(oddnums...)
}