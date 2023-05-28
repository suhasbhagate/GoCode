package main

import(
	"fmt"
	"unicode"
)


func main(){
	str:="Suhas!1"
	var flgnumber, flglower, flgupper, flgspecial bool
	for _, v :=  range str{
		switch {
		case unicode.IsUpper(v):
			flgupper = true
		case unicode.IsLower(v):
			flglower = true
		case unicode.IsNumber(v):
			flgnumber = true
		case unicode.IsSymbol(v) || unicode.IsPunct(v):
			flgspecial = true
		}
	}
	if !flgupper{
		fmt.Println("Missing Uppercase Letter")
	}
	if !flglower{
		fmt.Println("Missing Lowercase Letter")
	}
	if !flgnumber{
		fmt.Println("Missing Number")
	}
	if !flgspecial{
		fmt.Println("Missing Special Character")
	}
	if flglower && flgupper && flgnumber && flgspecial{
		fmt.Println("Strong Password")
	}
}