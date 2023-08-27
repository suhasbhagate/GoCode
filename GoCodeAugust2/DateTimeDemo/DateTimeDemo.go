package main

import(
	"fmt"
	"time"
)

func main(){
	//fmt.Println(time.Now().GoString())
	m,_ := time.ParseDuration("11h30m")
	then := time.Now().Add(time.Duration(m.Minutes()))
	fmt.Println(then)
}