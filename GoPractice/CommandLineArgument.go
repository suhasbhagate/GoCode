package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Program Name ",os.Args[0])
	fmt.Println("Arguments ", os.Args[1:])
}