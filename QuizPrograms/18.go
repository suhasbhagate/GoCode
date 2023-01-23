package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x [100000]struct{}
	fmt.Println(unsafe.Sizeof(x))
}