package main
import (
"fmt"
"time"
) 

func main() {
go runGoroutine()
fmt.Println("Main begun")
time.Sleep(5 * time.Second)
fmt.Println("Finished Main")
} 

func runGoroutine() {
fmt.Println("Inside Goroutine")
}