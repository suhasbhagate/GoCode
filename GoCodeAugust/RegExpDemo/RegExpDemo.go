package main

import(
	"fmt"
	"regexp"
)

func main(){
	str1 := "abcd13efghi3jkl6mno2pq4rs0tuv8wxyz0"

	str2 := regexp.MustCompile(`[^a-zA-Z]`).ReplaceAllString(str1,"")
	fmt.Println(str2)
}