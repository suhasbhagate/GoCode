package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "github.com/sbbhagate/GoCode/Employee/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var collection *mongo.Collection
var collection1 *mongo.Collection
var grpcaddr string = "0.0.0.0:50052"
var httpaddr string = "0.0.0.0:50051"

type Server struct {
	pb.EmployeeServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://m001-student:m001password@sandbox.lsj2y.mongodb.net/test"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("EmployeeDB").Collection("Employee")
	collection1 = client.Database("EmployeeDB").Collection("emp_mat_view")
	//RunGRPCServer()
	RunGatewayServer()
}

func RunGRPCServer() {
	lis, err := net.Listen("tcp", grpcaddr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", grpcaddr)

	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failded to serve: %v \n", err)
	}
}

func RunGatewayServer() {
	grpcMux := runtime.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := pb.RegisterEmployeeServiceHandlerServer(ctx, grpcMux, &Server{})
	if err != nil {
		log.Fatal("Cannot register handler server:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	lis, err := net.Listen("tcp", httpaddr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("HTTP GateWay Server started at %s\n", httpaddr)

	err = http.Serve(lis, mux)
	if err != nil {
		log.Fatalf("Failded to start HTTP Gateway Server: %v \n", err)
	}
}
