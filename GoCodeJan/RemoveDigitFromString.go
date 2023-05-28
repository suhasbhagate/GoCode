package main

import(
	"fmt"
	"regexp"
)

func main(){
	str := "G7ol56ang12Pro8gr09ams34"
	str = regexp.MustCompile(`[0-9]`).ReplaceAllString(str, "")
	fmt.Println(str)
}