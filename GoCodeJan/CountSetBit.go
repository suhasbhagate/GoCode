package main

import(
	"fmt"
)

func main(){
	fmt.Println(CountSetBit(255))
}

func CountSetBit(num int)int{
	cnt := 0
	for num != 0{
		if (num & 1) >0{
			cnt ++
		}
		num >>= 1
	}
	return cnt
}