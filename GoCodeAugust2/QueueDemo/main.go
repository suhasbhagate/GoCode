package main

import(
	"fmt"
	"example.com/demo/QueueDemo/queue"
)

func main(){
	var qu queue.Queue
	qu.Enqueue(100)
	qu.Enqueue(200)

	fmt.Println(qu)
	if !qu.IsEmpty(){
		fmt.Println(qu.Dequeue())
		fmt.Println(qu)
	}
	if !qu.IsEmpty(){
		fmt.Println(qu.Dequeue())
		fmt.Println(qu)
	}
	if !qu.IsEmpty(){
		fmt.Println(qu.Dequeue())
		fmt.Println(qu)
	}
}