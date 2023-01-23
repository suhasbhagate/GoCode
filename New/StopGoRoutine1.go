package main

import (
	"fmt"
)

func main(){
	ch := make(chan string)

	go func(){
		for{
			v, ok:= <-ch
			if !ok{
				fmt.Println("Completed Receiving from ch")
				return
			}
			fmt.Printf("Received %v\n", v)
		}
	}()

	// go func(){
	// 	for v := range ch{
	// 		fmt.Printf("Received %v\n",v)
	// 	}
	// }()

	ch <- "Hi"
	ch <- "Hello"
	ch <- "World"
	//close(ch)
	fmt.Println("Finished")
}