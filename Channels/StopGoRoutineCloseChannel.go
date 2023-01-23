package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make (chan string)

	go func(){
		// for{
		// 	v, ok:= <-ch
		// 	if !ok {
		// 		fmt.Println("Finished")
		// 		return
		// 	}
		// 	fmt.Println(v)
		// }

		for v := range ch{
			fmt.Println(v)
		}
	}()

	ch <- "Hi"
	ch <- "Hello"
	ch <- "Good Morning"
	close(ch)
	time.Sleep(time.Second)
}