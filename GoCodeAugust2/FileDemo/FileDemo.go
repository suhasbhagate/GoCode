package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("D:\\GoCode\\GoCodeAugust2\\FileDemo\\demo.txt")
	if err != nil {
		log.Println(err)
	}
	file.WriteString("Saksham Suhas Bhagate")
	data, err := ioutil.ReadFile("D:\\GoCode\\GoCodeAugust2\\FileDemo\\demo.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))
}
