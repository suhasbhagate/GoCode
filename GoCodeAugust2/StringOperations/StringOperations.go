package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main(){
	strset := "Saksham Suhas Bhagate Nandani Suhas Bapuso Bhagate Nandani"
	strcheck := "Saksham"

	fmt.Println(strings.Contains(strset,strcheck))
	xs := strings.Fields(strset)
	fmt.Println(xs)

	mp := make(map[string]int)
	for _,v := range xs{
		mp[v]++
	}
	fmt.Println(mp)

	strnew := strings.Join(xs," ")
	fmt.Println(strnew)

	xsnew := strings.Split(strnew," ")
	fmt.Println(xsnew)

	var sb strings.Builder
	sb.WriteString("Hello")
	sb.WriteString(" Saksham")
	fmt.Println(sb.String())

	line := "At Post Nandani Taluka Shirol District Kolhapur"
	before, after, found := strings.Cut(line,"Nandani")
	fmt.Println(before)
	fmt.Println(after)
	fmt.Println(found)

	line2 := "ABCDEF!1GHI@#JKL$MNK5QUI%5JK^6MFK&7kfo*8aod(9)"

	str := regexp.MustCompile(`[^A-Za-z0-9]`).ReplaceAllString(line2,"")
	fmt.Println(str)

	fmt.Println(VerifyPassword("Sa1#"))

}

func VerifyPassword(password string)bool{
	var flgUP, flgLW, flgNUM, flgSPEC bool
	for _,v := range password{
		switch {
		case unicode.IsUpper(v):
			flgUP = true

		case unicode.IsLower(v):
			flgLW = true

		case unicode.IsNumber(v):
			flgNUM = true

		case unicode.IsPunct(v) || unicode.IsSymbol(v):
			flgSPEC = true
		}		
	}
	if !flgUP{
		fmt.Println("Missing Upper Case Letter")
	}
	if !flgLW{
		fmt.Println("Missing Lower Case Letter")
	}
	if !flgNUM{
		fmt.Println("Missing Number")
	}
	if !flgSPEC{
		fmt.Println("Missing Special Character")
	}
	return flgUP && flgLW && flgNUM && flgSPEC
}

