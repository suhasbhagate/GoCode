package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	pb "github.com/sbbhagate/GoCode/Employee/proto"
	sng "github.com/sbbhagate/GoCode/Employee/singleton"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var grpcaddr string
var httpaddr string

type Server struct {
	pb.EmployeeServiceServer
}

func init(){
	err := godotenv.Load("local.env")
    if err != nil {
        sng.SngService.Fatal("Error loading .env file")
    }
}

func main() {	
	
	grpcaddr = os.Getenv("GRPCADDR")
	httpaddr = os.Getenv("HTTPADDR")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sng.InitSingletonLogger(ctx)
	defer sng.SngService.Sync()
	
	sng.SngService.Debug("Employee Service API Started")
	sng.InitMongoDB(ctx)

	defer sng.MongoService.DisconnectMongoDB(ctx)

	RunGatewayServer(ctx)
}

func RunGRPCServer() {
	lis, err := net.Listen("tcp", grpcaddr)
	if err != nil {
		sng.SngService.Fatal("Failed to listen on: %v\n", zap.Error(err))
	}

	sng.SngService.Debug("Listening on ", zap.String(" %s\n", grpcaddr))

	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		sng.SngService.Fatal("Failded to serve: %v \n", zap.Error(err))
	}
}

func RunGatewayServer(ctx context.Context) {
	grpcMux := runtime.NewServeMux()

	err := pb.RegisterEmployeeServiceHandlerServer(ctx, grpcMux, &Server{})
	if err != nil {
		sng.SngService.Fatal("Cannot register handler server:", zap.Error(err))
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	lis, err := net.Listen("tcp", httpaddr)
	if err != nil {
		sng.SngService.Fatal("Failed to listen on: %v\n", zap.Error(err))
	}

	str := fmt.Sprintf("HTTP GateWay Server started at %s \n",httpaddr)
	sng.SngService.Debug(str)

	err = http.Serve(lis, mux)
	if err != nil {
		sng.SngService.Fatal("Failded to start HTTP Gateway Server: %v \n", zap.Error(err))
	}
}