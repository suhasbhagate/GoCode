package main

import(
	"fmt"
)

func main(){
	ch := make(chan string)
	go func(){
		// for{
		// 	v,ok := <-ch
		// 	if !ok{
		// 		fmt.Println("Stopped Receiving")
		// 		return
		// 	}
		// 	if ok{
		// 		fmt.Println("Received ",v)
		// 	}
		// }
		for v:= range ch{
			fmt.Println("Received ",v)
		}
	}()
	ch <- "Hi"
	ch <- "Hello"
	ch <- "GM"
	close(ch)
	fmt.Println("Exit")
}