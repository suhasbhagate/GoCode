package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Results struct{
	Sum int
	Product int
}

//Function to add number to Sum (Field of struct Results)
func (r *Results) Add(wg *sync.WaitGroup,num int){ 
	defer wg.Done()
	r.Sum += num
}

//Function to multiply number to Product (Field of struct Results)
func (r *Results) Multiply(wg *sync.WaitGroup,num int){
	defer wg.Done()
	r.Product *= num
}

//Function to display Sum and Product field of struct Results
func (r *Results) Display(){
	fmt.Printf("Sum: %v\t Product: %v\n",r.Sum,r.Product)
}

func main(){	
	r := Results{
		Sum: 0,
		Product: 1,
	}
	//wg.Add(40)
	for i:=1; i<=20; i++{
		wg.Add(2)				
		go r.Add(&wg, i)
		go r.Multiply(&wg, i)
	}
	wg.Wait()
	r.Display()
}