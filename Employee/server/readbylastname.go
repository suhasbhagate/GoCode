package main

import (
	"context"
	"log"

	pb "github.com/sbbhagate/GoCode/Employee/proto"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadEmployeeByLastName(ctx context.Context, in *pb.EmployeeLastName) (*pb.Employees, error) {
	log.Printf("ReadEmployeeByLastName is invoked with %v\n", in)

	pipeline, err := collection.Aggregate(ctx, bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "lastname", Value: in.LastName}}}},
	})
	if err != nil {
		log.Fatal(err)
	}

	defer pipeline.Close(ctx)

	emplist := &pb.Employees{}

	for pipeline.Next(ctx) {
		data := &Emp{}
		err := pipeline.Decode(data)

		if err != nil {
			return nil, status.Errorf(
				codes.NotFound,
				err.Error(),
			)
		}
		emplist.Emps = append(emplist.Emps, documentToEmployee(data))
	}
	if err := pipeline.Err(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return emplist, nil
}
