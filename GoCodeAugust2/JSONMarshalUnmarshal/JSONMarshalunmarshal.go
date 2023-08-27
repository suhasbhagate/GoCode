package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string
	LastName  string
}

var people1 = []Person{{"Saksham", "Bhagate"}, {"Suhas", "Bhagate"}}
var people2 []Person

func main() {
	bs1, err := json.Marshal(people1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs1))

	bs2, err := json.MarshalIndent(people1,"","\t")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs2))

	err = json.Unmarshal(bs2,&people2)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(people2)

	err = json.Unmarshal(bs1,&people2)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(people2)
}
