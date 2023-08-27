package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ctx,cancel := context.WithCancel(context.Background())
	ch := make(chan string)

	go func(ctx context.Context){
		for {
			select{
			case <- ctx.Done():
				fmt.Println("Stopping Go Routine")
				close(ch)
				return

			default:
				ch <- "Hello"
			}
			time.Sleep(100 * time.Millisecond)
		}
	}(ctx)

	go func(ctx context.Context){
		time.Sleep(2 * time.Second)
		cancel()
	}(ctx)

	for v:= range ch{
		fmt.Println(v)
	}
}