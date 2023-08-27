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
	p1 := Person{
		FirstName: "Saksham",
		LastName: "Bhagate",
		Age: 15,
		Address: Address{City: "Nandani", Pin: 416102},
	}
	fmt.Println(p1)

	p2 := &Person{
		FirstName: "Suhas",
		LastName: "Bhagate",
		Age: 15,
		Address: Address{City: "Nandani", Pin: 416102},
	}
	fmt.Println(*p2)

	p3 := new (Person)
	p3 = p2
	fmt.Println(*p3)

	e := Employee(p1)
	fmt.Println(e)
}