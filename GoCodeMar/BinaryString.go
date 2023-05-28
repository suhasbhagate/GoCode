package main

import(
	"fmt"
	"strconv"
)

func main(){
	str := "1110"
	num,_:=strconv.ParseInt(str,2,64)
	fmt.Println(num)
	st := strconv.FormatInt(14,2)
	fmt.Println(st)
}