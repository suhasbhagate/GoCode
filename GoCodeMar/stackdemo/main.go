package main

import(
	"fmt"
	"example.com/gocodemar/stackdemo/stack"
)

func main(){
	var s stack.Stack
	s.Push(9)
	s.Push(10)
	s.Push(11)

	//fmt.Println(s.items)

	if !s.IsEmpty(){
		fmt.Println(s.Pop())
	}
	if !s.IsEmpty(){
		fmt.Println(s.Pop())
	}
	if !s.IsEmpty(){
		fmt.Println(s.Pop())
	}
	if !s.IsEmpty(){
		fmt.Println(s.Pop())
	}
}