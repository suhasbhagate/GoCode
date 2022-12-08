package main

import (
	"context"
	"fmt"
	pb "github.com/sbbhagate/GoCode/Employee/proto"
	sng "github.com/sbbhagate/GoCode/Employee/singleton"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadEmployeeByDesignation(ctx context.Context, in *pb.EmployeeDesignation) (*pb.Employees, error) {
	str := fmt.Sprintf("ReadEmployeeByDesignation is invoked with %v\n", in)
	sng.SngService.Debug(str)

	pipeline, err := sng.MongoService.Collection.Aggregate(ctx, bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "designation", Value: in.Designation}}}},
	})
	if err != nil {
		sng.SngService.Fatal("Error while searching data")
	}

	defer pipeline.Close(ctx)

	emplist := &pb.Employees{}

	for pipeline.Next(ctx) {
		data := &Emp{}
		err := pipeline.Decode(data)

		if err != nil {
			sng.SngService.Error("Error while retrieving data")
			return nil, status.Errorf(
				codes.NotFound,
				err.Error(),
			)
		}
		emplist.Emps = append(emplist.Emps, documentToEmployee(data))
	}
	if err := pipeline.Err(); err != nil {
		sng.SngService.Error("Error while retrieving data")
		return nil, status.Error(codes.Internal, err.Error())
	}
	return emplist, nil
}