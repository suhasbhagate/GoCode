package main

import (
	"context"
	"fmt"
	"time"
)


func main(){
	ch := make(chan string)
	ctx,cancel := context.WithCancel(context.Background())
	go func(ctx context.Context){
		for{
			select{
			case <-ctx.Done():
				fmt.Println("Received Cancellation Signal")
				close(ch)
				return
			default:
				ch<-"Hi"
			}
		}
		time.Sleep(100 * time.Millisecond)
	}(ctx)

	go func(){
		time.Sleep(2 * time.Second)
		cancel()
	}()

	for v:=range ch{
		fmt.Println(v)
	}

}