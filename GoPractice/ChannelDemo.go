package main

import (
	"fmt"
)

func main(){
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go Send(even,odd,quit)
	Receive(even,odd, quit)
	fmt.Println("Completed")
}

func Send(e, o, q chan<- int){
	for i:=0; i<=100;i++{
		if i%2 ==0{
			e<- i
		} else{
			o <- i
		}
	}
	q <- 0
}

func Receive(e, o, q <-chan int){
	for{
		select{
		case v:= <-e:
			fmt.Printf("Even Channel Received %v\n",v)
		case v:= <-o:
			fmt.Printf("Odd Channel Received %v\n",v)
		case v := <-q:
			fmt.Printf("Quite Channel Received %v\n",v)
			return
		}
	}

}
