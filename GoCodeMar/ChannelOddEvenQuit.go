package main

import(
	"fmt"
)

func main(){
	odd := make(chan int)
	even := make(chan int)
	quit := make(chan int)

	go Send(odd,even,quit)
	Receive(odd,even,quit)
}


func Send(o, e, q chan<- int){
	for i:=1; i<100; i++{
		if i%2 ==0{
			e <- i
		} else{
			o <- i
		}
	}
	q <- 0
}

func Receive(o, e, q <-chan int){
	for{
		select{
		case v:= <-e:
			fmt.Println("Received from Even Channel ", v)
		case v := <-o:
			fmt.Println("Received from Odd Channel ", v)
		case v := <-q:
			fmt.Println("Received from Quit Channel ",v)
			return
		}
	}
}