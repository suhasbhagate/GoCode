package main

import (
	"log"
	"net"

	"github.com/suhasbhagate/GoCode/GoCodeAugust/GRPCDemo/pb"
	"github.com/suhasbhagate/GoCode/GoCodeAugust/GRPCDemo/handler"
	"google.golang.org/grpc"
)


func main() {
	listener, err := net.Listen("tcp", ":8181")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, &handler.Server{})
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}