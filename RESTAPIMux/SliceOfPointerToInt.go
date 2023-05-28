package main

import(
	"fmt"
)

func main() {
	pool := []int{1, 2, 3, 4, 5, 6}
	even := []*int{}
	
	for i := 0; i < len(pool); i++ {
		if pool[i]%2 == 0 {
		  even = append(even, &pool[i])
		}
	}
	
	for i := range even {
	  fmt.Print(*even[i], " ")
	}
}