package main

import(
	"fmt"
)

func main(){
	ch := make(chan string)
	// go func(){
	// 	for v := range ch{
	// 		fmt.Println(v)
	// 	} 
	// }()

	go func(){
		for{
			v,ok := <-ch
			if ok{
				fmt.Println(v)
			}
			if !ok{
				fmt.Println("Closed Channel")
				return
			}
		}
	}()

	ch <- "Hello"
	ch <- "Hi"
	//close(ch)
}