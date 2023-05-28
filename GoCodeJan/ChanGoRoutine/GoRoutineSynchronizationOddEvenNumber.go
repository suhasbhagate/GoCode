package main

import(
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	for i:= 1; i<=100; i++{
		if i%2 ==0{
			wg.Add(1)
			go func(i int){
				fmt.Println("GoRoutine 1: ", i)
				defer wg.Done()
			}(i)
			wg.Wait()
		} else{
			wg.Add(1)
			go func(i int){
				fmt.Println("GoRoutine 2: ", i)
				defer wg.Done()
			}(i)
			wg.Wait()
		}
	}
}