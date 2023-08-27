package main

import (
	"fmt"
)

func main() {
	const (
		_  = iota
		kb = 1 << (iota * 10)
		mb
		gb
	)
	fmt.Printf("KB: %d %b %o %x\n", kb, kb, kb, kb)
	fmt.Printf("MB: %d %b %o %x\n", mb, mb, mb, mb)
	fmt.Printf("GB: %d %b %o %x\n", gb, gb, gb, gb)
}
