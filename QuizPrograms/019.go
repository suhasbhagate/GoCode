package main
import (
	"fmt"
	"sync"
)
func main() {
	var wg sync.WaitGroup
	nums := make(chan int, 3)
	wg.Add(1)
	go func() {
		defer wg.Done()
		nums <- 10
		nums <- 30
		nums <- 50
		//close(nums)
	}()

	for v := range nums {
		fmt.Println(v)
	}
	wg.Wait()
}