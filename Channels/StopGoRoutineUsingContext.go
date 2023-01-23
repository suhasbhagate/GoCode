package main

import (
	"fmt"
	"time"
	"context"
)

func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
   
	go func(ctx context.Context) {
		for {
			select {
				case <-ctx.Done():
					ch <- struct{}{}
	   				return
	  			default:
	   				fmt.Println("Hello")
	  		}
   
	  	time.Sleep(500 * time.Millisecond)
	 	}
	}(ctx)
   
	go func() {
	 	time.Sleep(3 * time.Second)
	 	cancel()
	}()
   
	fmt.Println(<-ch)
	fmt.Println("Finish")
}