package main

import (
	"fmt"
	"sync"
)

func Hello(wg *sync.WaitGroup) {
	fmt.Println("Hello")
	defer wg.Done()
}

func World(wg *sync.WaitGroup) {
	fmt.Println("World")
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("I Say, ")

	go World(&wg)
	go Hello(&wg)
	wg.Wait()
}