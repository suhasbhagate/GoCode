package main

import (
	"fmt"
	"log"
	"net"
)

func main(){
	listener, err := net.Listen("tcp","localhost:2001")
	if err != nil{
		log.Println(err)
	}
	defer listener.Close()
	for{
		conn, err := listener.Accept()
		if err != nil{
			log.Println(err)
		}
		go HandleRequest(conn)
	}
}

func HandleRequest(conn net.Conn){
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil{
		log.Println(err)
	}
	str := fmt.Sprintf("Hello %v",string(buff[:]))
	_, err = conn.Write([]byte(str))
	if err != nil{
		log.Println(err)
	}
	conn.Close()
}