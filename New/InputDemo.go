package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, err := reader.ReadString('\n')
	//text, _, err := reader.ReadLine()
	if err!=nil{
		fmt.Println("Wrong Input")
	}
    fmt.Println(string(text))
}