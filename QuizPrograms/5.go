package main

import ("fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
		defer wg.Done()
		mu.Lock()
		m[i] = i
		mu.Unlock()
	}()
	}
	wg.Wait()
	fmt.Println(len(m))
}