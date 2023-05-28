package main

import (
	"fmt"
	//"time"
)

func main(){
	ch := make(chan string)
	
	go func(){
		// for v := range ch{
		// 	fmt.Println("Received from Ch ",v)
		// }

		for{
			v, ok := <-ch
			if !ok{
				fmt.Println("Stopped Receiving from Channel ch")
				return
			}
			fmt.Println("Received from Channel ch ",v)
		}
	}()

	ch <- "Hi"
	ch <- "Hello"
	ch <- "Good Morning"
	//close(ch)
	//time.Sleep(time.Second)
}