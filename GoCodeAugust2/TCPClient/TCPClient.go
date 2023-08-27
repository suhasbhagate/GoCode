package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr("tcp", "localhost:2001")
	if err != nil {
		log.Println(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpServer)
	if err != nil {
		log.Println(err)
	}

	_, err = conn.Write([]byte("Saksham Bhagate"))
	if err != nil {
		log.Println(err)
	}

	buff := make([]byte, 1024)
	_, err = conn.Read(buff)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(buff))
}
