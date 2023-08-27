package main

import (
	"fmt"
)

func main() {
	TypeSwitch(10)
	TypeSwitch(10.10)
	TypeSwitch("SA")
	TypeSwitch(true)
}

func TypeSwitch(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	default:
		fmt.Println("Unknown")
	}
}
