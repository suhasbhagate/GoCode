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
	fmt.Printf("Type of Channel %T Name of Channel %v\n",o,o)
	fmt.Printf("Type of Channel %T Name of Channel %v\n",e,e)
	fmt.Printf("Type of Channel %T Name of Channel %v\n",q,q)
	for i:= 1; i<=100; i++{
		if i%2 ==0{
			e <- i
		} else{
			o <- i
		}
	}
	q <- 0
}

func Receive(o, e, q <-chan int){
	fmt.Printf("Type of Channel %T Name of Channel %v\n",o,o)
	fmt.Printf("Type of Channel %T Name of Channel %v\n",e,e)
	fmt.Printf("Type of Channel %T Name of Channel %v\n",q,q)
	for{
		select{
		case v:= <-e:
			fmt.Printf("Received %v from Even Channel\n",v)
		case v:= <-o:
			fmt.Printf("Received %v from Odd Channel\n",v)
		case <-q:
			fmt.Println("Received Stop Signal from Quit Channel")
			return
		}
	}
}