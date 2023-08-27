package main

import (
	"fmt"

	"github.com/suhasbhagate/GoCode/GoCodeAugust/StackDemo/stack"
)

func main() {
	var st stack.Stack
	st.Push(10)
	st.Push(20)
	st.Push(30)
	fmt.Println(st)

	fmt.Println("Top: ", st.Top())

	fmt.Println(st.IsEmpty())

	if !st.IsEmpty() {
		fmt.Println(st.Pop())
	}
	if !st.IsEmpty() {
		fmt.Println(st.Pop())
	}
	if !st.IsEmpty() {
		fmt.Println(st.Pop())
	}
	if !st.IsEmpty() {
		fmt.Println(st.Pop())
	}
	if !st.IsEmpty() {
		fmt.Println(st.Pop())
	}

	fmt.Println("Top ", st.Top())
}
