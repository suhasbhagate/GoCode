package main

import (
	"fmt"
	"sort"
)

type Person struct{
	FirstName string
	LastName string
	Age int
}

type ByAge []Person

func (b ByAge) Len()int{
	return len(b)
}

func (b ByAge) Less(i, j int)bool{
	return b[i].Age < b[j].Age
}

func (b ByAge) Swap(i, j int){
	b[i],b[j] = b[j], b[i]
}

func main(){
	p1 := Person{"Suhas", "Bhagate", 43}
	p2 := Person{"Saksham", "Bhagate", 15}
	p3 := Person{"Atharva", "Bhagate", 18}

	people := []Person{p1, p2, p3}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}