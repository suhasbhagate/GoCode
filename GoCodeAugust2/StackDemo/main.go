package main

import(
	"example.com/demo/StackDemo/stack"
	"fmt"
)

func main(){
	var st stack.Stack
	st.Push(1001)
	st.Push(1002)
	fmt.Println(st)

	if !st.IsEmpty(){
		fmt.Println(st.Pop())
		fmt.Println(st)
	}
	if !st.IsEmpty(){
		fmt.Println(st.Pop())
		fmt.Println(st)
	}
	if !st.IsEmpty(){
		fmt.Println(st.Pop())
		fmt.Println(st)
	}
}