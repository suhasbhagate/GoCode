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
				fmt.Println("Received Signal to Stop")
				close(ch)
				return
			default:
				ch <- "Hi"
				time.Sleep(100 * time.Microsecond)
			}
		}
	}()

	go func(){
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	for v := range ch{
		fmt.Println(v)
	}
}