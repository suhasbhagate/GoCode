package main

import(
	"fmt"
)

func main(){
	// chan1 := make(chan int)
	// go func(chan1 chan int){
	// 	chan1 <- 10
	// }(chan1)
	// fmt.Println(<-chan1)

	chan2 := make(chan int, 2)
	chan2 <- 10
	chan2 <- 20
	//chan2 <- 30
	fmt.Println(<-chan2)
	fmt.Println(<-chan2)	
	//fmt.Println(<-chan2)

}