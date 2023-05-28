package main

import(
	"fmt"
	"example.com/gocodejan/StackDemo/Stack"
)

func main(){
	var st stack.Stack
	st.Push(101)
	st.Push(102)
	st.Push(103)
	fmt.Println(st)

	if !st.IsEmpty(){
		fmt.Println("Top",st.Top())
		fmt.Println("Popped ",st.Pop())
	}
	if !st.IsEmpty(){
		fmt.Println("Top",st.Top())
		fmt.Println("Popped ",st.Pop())
	}
	if !st.IsEmpty(){
		fmt.Println("Top",st.Top())
		fmt.Println("Popped ",st.Pop())
	}
	if !st.IsEmpty(){
		fmt.Println("Top",st.Top())
		fmt.Println("Popped ",st.Pop())
	}
}