package main

import (
	"fmt"
	"time"
)

func main(){
	ch:= make(chan string)
	done:= make(chan struct{})

	go func(){
		for{
		select{
		case <-done:
			fmt.Println("Received Signal from done channel to stop")
			close(ch)
			return
		default:
			ch<-"Hi"
		}
		time.Sleep(100 * time.Millisecond)
		}
	}()

	go func(){
		time.Sleep(2 * time.Second)
		done<-struct{}{}
	}()

	for v:= range ch{
		fmt.Println("Received: ",v)
	}
	fmt.Println("Finish")
}