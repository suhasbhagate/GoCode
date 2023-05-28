package main

import(
	"fmt"
	s "example.com/gocodefeb/stackdemo/stack"
)

func main(){
	var st s.Stack
	st.Push(50)
	st.Push(60)

	fmt.Println("Top: ", st.Peek())
	if !st.IsEmpty(){
		fmt.Println("Popped: ",st.Pop())
	}
	if !st.IsEmpty(){
		fmt.Println("Popped: ",st.Pop())
	}
	if !st.IsEmpty(){
		fmt.Println("Popped: ",st.Pop())
	}
}