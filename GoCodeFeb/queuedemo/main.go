package main

import(
	"fmt"
	qu "example.com/gocodefeb/queuedemo/queue"
)

func main(){
	var q qu.Queue

	q.Enque(70)
	q.Enque(80)
	fmt.Println("Front Element: ",q.Peek())

	if !q.IsEmpty(){
		fmt.Println("Element: ",q.Deque())
	}
	if !q.IsEmpty(){
		fmt.Println("Element: ",q.Deque())
	}
	if !q.IsEmpty(){
		fmt.Println("Element: ",q.Deque())
	}
}