package main

import (
	"context"
	"fmt"
	"strconv"
	pb "github.com/sbbhagate/GoCode/Employee/proto"
	sng "github.com/sbbhagate/GoCode/Employee/singleton"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *Server) ReadEmployeeById(ctx context.Context, in *pb.EmployeeId) (*pb.Employees, error) {
	str := fmt.Sprintf("ReadEmployeeById is invoked with %v\n", in)
	sng.SngService.Debug(str)

	pipeline, err := sng.MongoService.MatView.Aggregate(ctx, 
						bson.A{
							bson.D{
								{Key: "$search",
									Value: bson.D{
										{Key: "index", Value: "EmpIdIndex"},
										{Key: "text",
											Value: bson.D{
												{Key: "query", Value: strconv.Itoa(int(in.Empid))},
												{Key: "path", Value: "empid"},
											},
										},
									},
								},
							},
						})

	if err != nil {
		sng.SngService.Fatal("Error while searching data")
	}

	defer pipeline.Close(ctx)

	emplist := &pb.Employees{}

	for pipeline.Next(ctx) {
		data := &Emp1{}
		err := pipeline.Decode(data)

		if err != nil {
			sng.SngService.Error("Error while retrieving data")
			return nil, status.Errorf(
				codes.NotFound,
				err.Error(),
			)
		}
		emplist.Emps = append(emplist.Emps, documentToEmployee1(data))
	}
	if err := pipeline.Err(); err != nil {
		sng.SngService.Error("Error while retrieving data")
		return nil, status.Error(codes.Internal, err.Error())
	}
	return emplist, nil
}