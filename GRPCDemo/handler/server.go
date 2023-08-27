package handler

import(
	"github.com/suhasbhagate/GoCode/GoCodeAugust/GRPCDemo/pb"
)

type Server struct {
	pb.UnimplementedServiceServer
}