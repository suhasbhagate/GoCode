package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Create("D:\\GoCode\\GoCodeAugust\\FileDemo\\demo.txt")
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("Saksham Suhas Bhagate Nandani")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile("D:\\GoCode\\GoCodeAugust\\FileDemo\\demo.txt")
	fmt.Println(string(data))
}
