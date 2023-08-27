package main

import(
	"fmt"
)

func main(){
	odd := make (chan int)
	even := make(chan int)
	quit := make(chan int)

	go Send(odd, even, quit)
	Receive(odd, even, quit)
}

func Send (odd, even, quit chan<- int){
	for i:=1; i<=100; i++{
		if i%2 == 0{
			even <- i
		} else{
			odd <- i
		}
	}
	quit <- 0
}

func Receive(odd, even, quit <-chan int){
	for{
		select{
		case v:= <- odd:
			fmt.Println("Received from Odd Channel: ",v)
			
		case v:= <- even:
			fmt.Println("Received from Even Channel: ",v)

		case v := <- quit:
			fmt.Println("Received from Quit Channel: ",v)
			return
		}
	}
}