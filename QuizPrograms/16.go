package main

import (
	"fmt"
)

func main(){
	var m map[string]string
	//m := make(map[string]string)
	m["go"] = "Golang" // Verify this work or not

	fmt.Println(m)
}