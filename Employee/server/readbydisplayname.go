package main

import (
	"context"
	"fmt"
	pb "github.com/sbbhagate/GoCode/Employee/proto"
	sng "github.com/sbbhagate/GoCode/Employee/logger"
	empdb "github.com/sbbhagate/GoCode/Employee/db"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadEmployeeByDisplayName(ctx context.Context, in *pb.EmployeeDisplayName) (*pb.Employees, error) {
	str := fmt.Sprintf("ReadEmployeeByDisplatName is invoked with %v\n", in)
	sng.DebugLogger.Debug(ctx, str, data)

	pipeline, err := empdb.MongoService.Collection.Aggregate(ctx, 
						bson.A{
							bson.D{
								{Key: "$search",
									Value: bson.D{
										{Key: "index", Value: "DisplayNameIndex"},
										{Key: "text",
											Value: bson.D{
												{Key: "query", Value: in.DisplayName},
												{Key: "path", Value: "displayname"},
											},
										},
									},
								},
							},
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