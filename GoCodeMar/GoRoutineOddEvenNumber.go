package main

import(
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	for i:=0; i<=100;i++{
		if i%2 ==0{
			wg.Add(1)
			go func (i int){
				defer wg.Done()
				fmt.Println("First Goroutine ", i)
			}(i)
			wg.Wait()
		} else{
			wg.Add(1)
			go func(i int){
				defer wg.Done()
				fmt.Println("Second Goroutine ",i)
			}(i)
			wg.Wait()
		}
	}
}