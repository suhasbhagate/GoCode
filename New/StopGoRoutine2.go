package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make (chan string)
	done := make (chan struct{})

	go func(){
		for{
			select{
			case ch<-"Hello":

			case <-done:
				fmt.Println("Received signal from done channel")
				close(ch)
				return
			}
		}
		time.Sleep(100 * time.Millisecond)
	}()

	go func(){
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	for v:=range ch{
		fmt.Printf("Received Data %v\n", v)
	}
	fmt.Println("Finished")
}