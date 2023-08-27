package handler

import(
	"context"
	"example.com/demo/GRPCDemo/pb"
)

func (server *Server) GetUserService(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error){
	usid := req.GetUserId()
	var usr *pb.User
	for _,v := range userList{
		if usid == int32(v.UserId){
			usr = &pb.User{
				UserId: int32(v.UserId),
				UserName: v.UserName,
			}
		}
	}
	return &pb.GetUserResponse{User: usr}, nil
}