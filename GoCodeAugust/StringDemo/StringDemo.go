package main

import(
	"fmt"
	"strings"
)

func main(){
	str := "Saksham Suhas Bhagate Nandani Suhas Bapuso Bhagate Nandani"
	fmt.Println(CountOccurencesOfStrings(str))

	slc := strings.Fields(str)
	fmt.Println(slc)
	fmt.Println(JoinStrings(slc," "))
}

func CountOccurencesOfStrings(str string)map[string]int{
	slc := strings.Fields(str)
	mp := make(map[string]int)

	for _, v:= range slc{
		mp[v]++
	}
	return mp
}

func JoinStrings(slc []string, joiner string)string{
	str := strings.Join(slc, joiner)
	return str
}