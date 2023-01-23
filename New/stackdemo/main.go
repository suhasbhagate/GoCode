package main

import (
	"fmt"
	"example.com/demo/stackdemo/stack"
)

func main(){
	var st stack.Stack

	st.Push(1)
	st.Push(2)
	st.Push(3)

	fmt.Println(st.Top())
	fmt.Println(st.IsEmpty())

	if !st.IsEmpty(){
		fmt.Println(st.Pop())
	}
	if !st.IsEmpty(){
		fmt.Println(st.Pop())
	}
	if !st.IsEmpty(){
		fmt.Println(st.Pop())
	}
	if !st.IsEmpty(){
		fmt.Println(st.Pop())
	}
	fmt.Println(st.Top())
}