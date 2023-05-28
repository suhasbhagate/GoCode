package main

import (
	"fmt"
)

func main(){
	var i int32 = 1
	var j uint32
	j = uint32(^i)
	fmt.Println(int(j))
}