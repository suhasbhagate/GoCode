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
	"google.golang.org/grpc"
)

var grpcaddr string
var httpaddr string
var data interface{}

type Server struct {
	pb.EmployeeServiceServer
}

// func init(){
// 	err := godotenv.Load("local.env")
//     if err != nil {
//         sng.SngService.Fatal("Error loading .env file")
//     }
// }

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	err := godotenv.Load("local.env")
    if err != nil {
		sng.FatalLogger.Fatal(ctx,err,"Error loading .env file",data)
    }	
	
	grpcaddr = os.Getenv("GRPCADDR")
	httpaddr = os.Getenv("HTTPADDR")

    //fmt.Println(os.Getenv("GRPCADDR"))
    //fmt.Println(os.Getenv("HTTPADDR"))
		
	defer cancel()

	sng.InitSingletonLogger(ctx,"Employee")
	//defer sng.SngService.Sync()
	
	sng.DebugLogger.Debug(ctx, "Employee Service API Started", data)
	sng.InitMongoDB(ctx)

	defer sng.MongoService.DisconnectMongoDB(ctx)

	RunGatewayServer(ctx)
	//RunGRPCServer(ctx)
}

func RunGRPCServer(ctx context.Context) {
	lis, err := net.Listen("tcp", grpcaddr)
	if err != nil {
		sng.FatalLogger.Fatal(ctx,err,"Failed to listen on GRPC Server",data)
	}

	str := fmt.Sprintf("Listening on %s\n", grpcaddr)
	sng.DebugLogger.Debug(ctx, str, data)

	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		sng.FatalLogger.Fatal(ctx,err,"Failed to Serve on GRPC Server",data)
	}
}

func RunGatewayServer(ctx context.Context) {
	grpcMux := runtime.NewServeMux()

	err := pb.RegisterEmployeeServiceHandlerServer(ctx, grpcMux, &Server{})
	if err != nil {
		sng.FatalLogger.Fatal(ctx,err,"Cannot register handler server",data)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	lis, err := net.Listen("tcp", httpaddr)
	if err != nil {
		sng.FatalLogger.Fatal(ctx,err,"Failed to listen on Gateway Server",data)
	}

	str := fmt.Sprintf("HTTP GateWay Server started at %s \n",httpaddr)
	sng.DebugLogger.Debug(ctx, str, data)

	err = http.Serve(lis, mux)
	if err != nil {
		sng.FatalLogger.Fatal(ctx,err,"Failded to start HTTP Gateway Server",data)
	}
}