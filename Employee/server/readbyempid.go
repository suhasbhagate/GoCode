package main

import (
	"context"
	"log"
	"strconv"

	pb "github.com/sbbhagate/GoCode/Employee/proto"
	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *Server) ReadEmployeeById(ctx context.Context, in *pb.EmployeeId) (*pb.Employees, error) {
	log.Printf("ReadEmployeeById is invoked with %v\n", in)

	// pipeline, err := collection.Aggregate(ctx, bson.A{
	// 	bson.D{{Key: "$match", Value: bson.D{{Key: "empid", Value: in.Empid}}}},
	// })


	pipeline, err := collection1.Aggregate(ctx, 
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
		log.Fatal(err)
	}

	defer pipeline.Close(ctx)

	emplist := &pb.Employees{}

	for pipeline.Next(ctx) {
		data := &Emp1{}
		err := pipeline.Decode(data)

		if err != nil {
			return nil, status.Errorf(
				codes.NotFound,
				err.Error(),
			)
		}
		emplist.Emps = append(emplist.Emps, documentToEmployee1(data))
	}
	if err := pipeline.Err(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return emplist, nil
}