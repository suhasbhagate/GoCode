package handler

import(
	"example.com/demo/GRPCDemo/pb"
)
type Server struct{
	pb.UnimplementedServiceServer
}