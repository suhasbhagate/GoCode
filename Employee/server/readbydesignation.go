package main

import (
	"context"
	"fmt"
	pb "github.com/suhasbhagate/GoCode/Employee/proto"
	sng "github.com/suhasbhagate/GoCode/Employee/logger"
	empdb "github.com/suhasbhagate/GoCode/Employee/db"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadEmployeeByDesignation(ctx context.Context, in *pb.EmployeeDesignation) (*pb.Employees, error) {
	str := fmt.Sprintf("ReadEmployeeByDesignation is invoked with %v\n", in)
	sng.DebugLogger.Debug(ctx, str, data)

	pipeline, err := empdb.MongoService.Collection.Aggregate(ctx, bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "designation", Value: in.Designation}}}},
	})
	if err != nil {
		sng.FatalLogger.Fatal(ctx,err,"Error while searching data",data)
	}

	defer pipeline.Close(ctx)

	emplist := &pb.Employees{}

	for pipeline.Next(ctx) {
		data := &Emp{}
		err := pipeline.Decode(data)

		if err != nil {
			sng.ErrorLogger.Error(ctx,err,"Error while retrieving data",data)
			return nil, status.Errorf(
				codes.NotFound,
				err.Error(),
			)
		}
		emplist.Emps = append(emplist.Emps, documentToEmployee(data))
	}
	if err := pipeline.Err(); err != nil {
		sng.ErrorLogger.Error(ctx,err,"Error while retrieving data",data)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return emplist, nil
}