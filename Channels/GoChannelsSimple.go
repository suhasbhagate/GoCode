package main

import (
	"fmt"
	"time"
)

func HelloSender(c chan string) {
	c <- "Hello"
}

func HelloReceiver(c chan string) {
	for {
		message := <-c
		fmt.Println(message)
	}
}

func main() {
	var c chan string = make(chan string)
	go HelloSender(c)
	go HelloReceiver(c)
	time.Sleep(2)
}
