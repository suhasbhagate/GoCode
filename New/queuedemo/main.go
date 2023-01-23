package main

import(
	"fmt"
	"example.com/demo/queuedemo/queue"
)

func main(){
	var q queue.Queue
	q.Enqueue(11)
	q.Enqueue(12)
	q.Enqueue(13)
	fmt.Println(q)

	fmt.Println("Front Element ", q.Peek())
	if !q.IsEmpty(){
		fmt.Println("Dequeued ", q.Dequeue())
	}
	fmt.Println("Front Element ", q.Peek())
	if !q.IsEmpty(){
		fmt.Println("Dequeued ", q.Dequeue())
	}
	fmt.Println("Front Element ", q.Peek())
	if !q.IsEmpty(){
		fmt.Println("Dequeued ", q.Dequeue())
	}
	fmt.Println("Front Element ", q.Peek())
	if !q.IsEmpty(){
		fmt.Println("Dequeued ", q.Dequeue())
	}
	fmt.Println("Front Element ", q.Peek())

}