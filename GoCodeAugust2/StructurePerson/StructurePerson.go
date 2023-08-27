package main

import(
	"fmt"
)

type Address struct{
	City string
	Pin int
}

type Person struct{
	FirstName string
	LastName string
	Age int
	Address Address
}

type Employee struct{
	FirstName string
	LastName string
	Age int
	Address Address
}

func main(){
	a1 := Address{City: "Nandani", Pin:416102}
	p1 := Person{FirstName: "Saksham", LastName: "Bhagate", Age: 15, Address: a1}
	fmt.Println(p1)
	
	a2 := Address{"Jaysingpur",416101}
	p2 := Person{"Suhas", "Bhagate", 42,a2}
	fmt.Println(p2)

	a3 := &Address{"Sangli", 416416}
	p3 := &Person{"Vihan", "Bhokare", 3,*a3}
	fmt.Println(*p3)

	a4 := new(Address)
	a4.City="Kolhapur"
	a4.Pin = 416001

	p4 := new(Person)
	p4.FirstName = "aaa"
	p4.LastName = "bbb"
	p4.Age = 60
	p4.Address = *a4
	fmt.Println(*p4)

	e := Employee(p1)
	fmt.Println(e)
}