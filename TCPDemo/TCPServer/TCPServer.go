package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8282")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go HandleRequest(conn)
	}
}

func HandleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	msg := fmt.Sprintf("Hello %v", string(buffer[:]))
	conn.Write([]byte(msg))
	conn.Close()
}
