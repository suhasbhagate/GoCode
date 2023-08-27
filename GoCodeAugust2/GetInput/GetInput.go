package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// var i, j int
	// _,err := fmt.Scan(&i, &j)
	// if err!= nil{
	// 	log.Println(err)
	// }
	// fmt.Println("i", i)
	// fmt.Println("j", j)

	// var ch rune
	// _,err := fmt.Scanf("%c",&ch)
	// if err!= nil{
	// 	log.Println(err)
	// }
	// fmt.Printf("ch %c", ch)

	reader := bufio.NewReader(os.Stdin)
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		lines = append(lines, line)
	}
	fmt.Println(lines)
}
