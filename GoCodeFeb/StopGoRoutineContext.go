package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ctx,cancel:= context.WithCancel(context.Background());
	ch := make(chan string)

	go func(ctx context.Context){
		for{
			select{
			case <-ctx.Done():
				fmt.Println("Received Cancel Signal")
				close(ch)
				return
			default:
				ch<- "Hi"
			}
			time.Sleep(100 * time.Millisecond)
		}
	}(ctx)

	go func(){
		time.Sleep(2 * time.Second)
		cancel()
	}()

	for v:= range ch{
		fmt.Println("Received: ",v)
	}
	fmt.Println("Finished")
}