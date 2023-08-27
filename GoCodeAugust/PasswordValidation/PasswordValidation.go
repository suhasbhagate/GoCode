package main

import(
	"fmt"
	"unicode"
)

func main(){
	Validate("ABCD")
	Validate("Ab@3")
}

func Validate(str string){
	fmt.Println(str)
	var UpperFlag, LowerFlag, NumberFlag, SpecialFlag bool
	for _,v:= range str{
		switch{
		case unicode.IsUpper(v):
			UpperFlag = true

		case unicode.IsLower(v):
			LowerFlag = true

		case unicode.IsNumber(v):
			NumberFlag = true

		case unicode.IsSymbol(v) || unicode.IsPunct(v):
			SpecialFlag = true
		}
	}
	if !UpperFlag{
		fmt.Println("No Upper Case Letter")
	}
	if !LowerFlag{
		fmt.Println("No Lower Case Letter")
	}
	if !NumberFlag{
		fmt.Println("No Number")
	}
	if !SpecialFlag{
		fmt.Println("No Special Character")
	}
	if UpperFlag && LowerFlag && NumberFlag && SpecialFlag{
		fmt.Println("Valid Password")
	}
}