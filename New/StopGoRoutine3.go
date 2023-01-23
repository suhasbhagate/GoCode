package main

import(
	"fmt"
	"time"
	"context"
)

func main(){
	ch := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context){
		for{
			select{
			case <- ctx.Done():
				close(ch)
				return
			default:
				ch <- "Hello"
			}
			time.Sleep(100 * time.Millisecond)
		}
	}(ctx)

	go func(){
		time.Sleep(2 * time.Second)
		cancel()
	}()

	for v:= range ch{
		fmt.Printf("Received %v\n",v)
	}
	fmt.Println("Finished")
}