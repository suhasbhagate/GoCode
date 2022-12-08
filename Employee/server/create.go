package main

import (
	"context"
	"fmt"
	pb "github.com/sbbhagate/GoCode/Employee/proto"
	sng "github.com/sbbhagate/GoCode/Employee/singleton"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateEmployee(ctx context.Context, in *pb.Employee) (*pb.EmployeeId, error){
	str := fmt.Sprintf("CreateEmployee was invoked with %v\n", in)
	sng.SngService.Debug(str)

	data := Emp{
		EmpId:	in.EmpId, 
		FirstName: in.FirstName,
		LastName: in.LastName,
		DisplayName: in.DisplayName,
		Age: in.Age,
		Salary: in.Salary,
		Designation: in.Designation,
		Department: in.Department,
	}
	
	res, err:= sng.MongoService.Collection.InsertOne(ctx, data)

	if err != nil{
		sng.SngService.Error("Error while inserting data")
		return nil, status.Errorf(
		codes.Internal,
		fmt.Sprintf("Internal Error: %v\n", err),
		)
	}

	
	_, ok := res.InsertedID.(primitive.ObjectID)
	if !ok{
		sng.SngService.Error("Cannot convert to OID ", zap.Error(err))
		sng.SngService.Error("Cannot convert to OID")
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	CreateMatView(ctx)
	
	return &pb.EmployeeId{
			Empid: data.EmpId,
		}, nil
}