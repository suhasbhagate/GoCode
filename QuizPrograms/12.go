package main
import (
	"fmt"
	"sync"
	"time"
)
func main() {
	finish := make(chan struct{})
	var done sync.WaitGroup
	done.Add(1)
	go func() {
		defer done.Done()
		select {
			case <-time.After(5 * time.Second):
			fmt.Println("time out")
			case <-finish:
			fmt.Println("done")
			default:
			fmt.Println("default")
		}
	}()
	close(finish)
	done.Wait()
}