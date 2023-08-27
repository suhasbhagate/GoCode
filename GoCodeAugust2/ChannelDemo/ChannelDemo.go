package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	fmt.Println(<-ch1)
	ch1 <- 20
	fmt.Println(<-ch1)
	ch1 <- 10
	fmt.Println(<-ch1)
	ch1 <- 20
	fmt.Println(<-ch1)

	// ch1 <- 30
	// ch1 <- 40
	// fmt.Println(<-ch1)
	// fmt.Println(<-ch1)

	ch2 := make(chan int)

	go func() {
		ch2 <- 15
	}()
	fmt.Println(<-ch2)

	odd := make(chan int)
	even := make(chan int)
	quit := make(chan int)

	go Send(odd, even, quit)
	Receive(odd, even, quit)
}

func Send(odd, even, quit chan<- int) {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- 0
}

func Receive(odd, even, quit <-chan int) {
	for{
		select{
		case v:= <-even:
			fmt.Println("Received from Even Channel: ",v)

		case v:=<-odd:
			fmt.Println("Received from Odd Channel: ",v)

		case v:=<-quit:
			fmt.Println("Received from Quit Channel: ",v)
			return
		}
	}
}
