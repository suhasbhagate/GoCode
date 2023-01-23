package main

import (
	"fmt"
)

func main() {
	m := map[string]int{}
	m["k1"] = 100
	v := m["k2"]
	fmt.Println(v)
}