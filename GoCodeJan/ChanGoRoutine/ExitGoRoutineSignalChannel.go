package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan string)
	done := make(chan struct{})

	go func(){
		for{
			select{
			case <-done:
				fmt.Println("Stopping Go Routine")
				close(ch)
				return
			default:
				ch<-"Hello"
			}
			time.Sleep(100*time.Millisecond)
		}
	}()

	go func(){
		time.Sleep(2*time.Second)
		done<-struct{}{}
	}()

	for v:= range ch{
		fmt.Println("Received ",v)
	}
	fmt.Println("Finished")
}