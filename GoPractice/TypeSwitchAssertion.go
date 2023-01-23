package main

import (
	"fmt"
)

func TypeSwitch(vr interface{}){
	switch vr.(type){
	case int:
		fmt.Printf("Type int : Value %v \n",vr.(int))

	case string:
		fmt.Printf("Type string : Value %v \n",vr.(string))

	case bool:
		fmt.Printf("Type bool : Value %v \n",vr.(bool))
	}
}

func main(){
	i:= 10
	j:="Saksham"
	k:=true
	TypeSwitch(i)
	TypeSwitch(j)
	TypeSwitch(k)
}