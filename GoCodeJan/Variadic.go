package main

import "fmt"

func foo(a ...int, b ...int) {
    fmt.Println(a,b)
}

func main() {
    a := 0
    //aa := 1
    b := []int{2,3,4}
    foo(a, b...)
}