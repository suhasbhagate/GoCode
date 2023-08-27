package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type ByAge []Person

var people = []Person{{"Saksham", "Bhagate", 15}, {"Suhas", "Bhagate", 43}, {"Atharva", "Bhagate", 15}}
var people2 = []Person{{"Saksham", "Bhagate", 15}, {"Suhas", "Bhagate", 43}, {"Atharva", "Bhagate", 15}}

func main() {
	intslc := []int{5, 2, 34, 86, 23}
	sort.Ints(intslc)
	fmt.Println(intslc)

	sort.Sort(sort.Reverse(sort.IntSlice(intslc)))
	fmt.Println(intslc)
	strslice := []string{"Saksham", "Suhas", "Bhagate", "Nandani"}
	sort.Strings(strslice)
	fmt.Println(strslice)

	sort.Slice(people, func(i, j int) bool {
		if people[i].Age != people[j].Age {
			return people[i].Age < people[j].Age
		}
		return people[i].FirstName < people[j].FirstName
	})
	fmt.Println(people)

	fmt.Println("Len Less Swap")
	sort.Sort(ByAge(people2))
	fmt.Println(people2)
}

func (b ByAge) Len()int{
	return len(b)
}

func (b ByAge) Less(i, j int)bool{
	return b[i].Age < b[j].Age
}

func (b ByAge) Swap(i, j int){
	b[i],b[j] = b[j],b[i]
}