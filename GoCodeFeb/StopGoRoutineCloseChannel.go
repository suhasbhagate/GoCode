package main

import(
	"fmt"
)

func main(){
	ch := make(chan string)
	go func(){
		for{
			v,ok:= <- ch
			if ok{
				fmt.Println("Received: ",v)
			} else{
				fmt.Println("Stopped Receiving")
				return
			}
		}
		// for v:= range ch{
		// 	fmt.Println("Received ",v)
		// }
	}()
	ch<- "Hi"
	ch<- "Hello"
	ch<- "Good Morning"
	//close(ch)
}