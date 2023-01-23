package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup

	for i:=1; i<=20; i++{
		if i % 2 == 0{
			wg.Add(1)
			go func(i int){
				fmt.Println("GoRoutine 1 Prints: ", i)
				defer wg.Done()
			}(i)
			wg.Wait()
		} else{
			wg.Add(1)
			go func(i int){
				fmt.Println("GoRoutine 2 Prints: ", i)
				defer wg.Done()
			}(i)
			wg.Wait()
		}
	}
}