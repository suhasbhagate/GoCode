package main

import(
	"fmt"
)

func Addition(a, b int)int{
	for b>0{
		carry := a & b
		a = a ^ b
		b = carry<<1
	}
	return a
}

func Subtraction (a, b int)int{
	for b!=0{
		borrow := ^a & b
		a = a ^ b
		b = borrow<<1
	}
	return a
}

func Double(a int)int{
	return a<<1
}

func Half(a int)int{
	return a>>1
}

func main(){
	a := 5
	b := 10
	resultAdd := Addition(a, b)
	fmt.Printf("Numbers %v + %v = %v\n",a,b,resultAdd)

	result := Subtraction(a, b)
	fmt.Printf("Numbers %v - %v = %v\n",a,b,result)

	fmt.Printf("Number %v Double %v Half %v\n",a,Double(a),Half(a))
}