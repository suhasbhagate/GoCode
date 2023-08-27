package main

import (
	"net"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr("tcp", "localhost:8282")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpServer)
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("Saksham Suhas Bhagate"))
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	println(string(buffer))
	conn.Close()
}
