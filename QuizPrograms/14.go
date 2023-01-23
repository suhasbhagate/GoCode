package main

import (
"fmt"
)

type employee struct {
	name string
	exp int
}

func (e employee) addExp() {
	e.exp++
}

func main() {
	e := employee{"rob", 10}
	e.addExp()
	fmt.Println(e.exp)
}