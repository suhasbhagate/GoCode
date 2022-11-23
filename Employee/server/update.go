package main

import (
	"context"
	"log"

	pb "github.com/sbbhagate/GoCode/Employee/proto"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateEmployee(ctx context.Context, in *pb.Employee) (*pb.UpdateEmployeeResponse, error) {
	log.Printf("UpdateEmployee was invoked with %v\n", in)

	data := Emp{
		EmpId:       in.EmpId,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		DisplayName: in.DisplayName,
		Age:         in.Age,
		Salary:      in.Salary,
		Designation: in.Designation,
		Department:  in.Department,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"empid": in.EmpId},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find Employee with given Id",
		)
	}

	CreateMatView(ctx)

	return &pb.UpdateEmployeeResponse{Response: "Updated Successfully"}, nil
}
