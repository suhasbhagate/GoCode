package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct{
	FirstName string
	LastName string
}

var people1,people2 []Person

func main(){
	people1 = []Person{{"Saksham","Bhagate"},{"Suhas","Bhagate"}}
	//bs, err := json.Marshal(people1)
	bs, err:= json.MarshalIndent(people1,"","\t")
	if err != nil{
		log.Println(err)
	}
	fmt.Println(string(bs))

	err = json.Unmarshal(bs,&people2)
	if err!= nil{
		log.Println(err)
	}
	fmt.Println(people2)

	json.NewEncoder(os.Stdout).Encode(people1)
}

