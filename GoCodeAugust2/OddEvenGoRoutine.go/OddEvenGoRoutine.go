package main

import(
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	for i:=1; i<=100; i++{
		if i%2 ==0{
			wg.Add(1)
			go func(wg *sync.WaitGroup, i int){
				fmt.Println("Even Channel: ",i)
				defer wg.Done()
			}(&wg,i)
			wg.Wait()
		} else{
			wg.Add(1)
			go func(wg *sync.WaitGroup, i int){
				fmt.Println("Odd Channel: ",i)
				defer wg.Done()
			}(&wg, i)
			wg.Wait()
		}
	}
}