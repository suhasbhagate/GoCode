package main

import(
	"fmt"
	"example.com/gocodejan/queuedemo/queue"
)

func main(){
	var q queue.Queue
	q.Enque(200)
	q.Enque("Saksham")
	q.Enque(10.10)
	fmt.Println(q)

	if !q.IsEmpty(){
		fmt.Println("Peek ",q.Peek())
		fmt.Println("Dequeue ",q.Deque())
	}
	if !q.IsEmpty(){
		fmt.Println("Peek ",q.Peek())
		fmt.Println("Dequeue ",q.Deque())
	}
	if !q.IsEmpty(){
		fmt.Println("Peek ",q.Peek())
		fmt.Println("Dequeue ",q.Deque())
	}
	if !q.IsEmpty(){
		fmt.Println("Peek ",q.Peek())
		fmt.Println("Dequeue ",q.Deque())
	}
	if !q.IsEmpty(){
		fmt.Println("Peek ",q.Peek())
		fmt.Println("Dequeue ",q.Deque())
	}
}

