package handler

import(
	"context"
	"example.com/demo/GRPCDemo/pb"
)

func (server *Server) AddUserService(ctx context.Context, req *pb.AddUserRequest) (*pb. AddUserResponse, error){
	usr := User{
		UserId: int(req.GetUser().UserId),
		UserName: req.GetUser().UserName,
	}
	userList = append(userList, usr)
	return &pb.AddUserResponse{Status: true}, nil
}
