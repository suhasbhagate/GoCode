package main

import(
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	for i:=0; i<100;i++{
		if i%2==0{
			wg.Add(1)
			go func(i int){
				defer wg.Done()
				fmt.Println("Go Routine 1: ", i)
			}(i)
			wg.Wait()
		} else{
			wg.Add(1)
			go func(i int){
				defer wg.Done()
				fmt.Println("Go Routine 2: ",i)
			}(i)
			wg.Wait()
		}
	}
}