package main

import (
	"fmt"
)

func main() {
	type mytype int
	var m mytype = 10
	fmt.Printf("Value: %v, Type: %T\n", m, m)
	n := int(m)
	fmt.Printf("Value: %v, Type: %T\n", n, n)
}
