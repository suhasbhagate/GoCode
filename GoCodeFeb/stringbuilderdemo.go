package main

import(
	"fmt"
	"strings"
)

func main(){
	slc := []string{"Saksham","Suhas","Bhagate","Nandani"}
	var sb strings.Builder
	for _,v := range slc{
		sb.WriteString(v)
		sb.WriteRune(' ')
		sb.WriteByte(' ')	
	}
	str:= strings.Join(slc," ")
	fmt.Println(sb.String())
	fmt.Println(str)
}