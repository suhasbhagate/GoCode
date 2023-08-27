package main

import(
	"fmt"
)

func main(){
	ch := make(chan string)

	// go func(ch chan string){
	// 	for v:= range ch{
	// 		fmt.Println("Value Received: ",v)
	// 	}
	// }(ch)

	go func(){
		for{
			v, ok := <-ch
			if !ok{
				fmt.Println("Stopped Receiving Value")
				return
			}
			if ok{
				fmt.Println("Received Value: ",v)
			}
		}
	}()

	ch <- "Hi"
	ch <- "Hello"
	//close(ch)
}