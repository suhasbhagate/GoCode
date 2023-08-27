package handler

import (
	"context"

	"github.com/suhasbhagate/GoCode/GoCodeAugust/GRPCDemo/pb"
)

func (server *Server) GetUserService(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var usr *pb.User
	userId := req.GetUserId()
	for _, v := range userList {
		if userId == int64(v.UserId) {
			usr = &pb.User{
				UserId:   int64(v.UserId),
				UserName: v.UserName,
			}
		}
	}
	return &pb.GetUserResponse{User: usr}, nil
}
