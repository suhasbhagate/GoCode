package main

import (
	"context"
	"fmt"
	pb "github.com/sbbhagate/GoCode/Employee/proto"
	sng "github.com/sbbhagate/GoCode/Employee/logger"
	empdb "github.com/sbbhagate/GoCode/Employee/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *Server) CreateEmployee(ctx context.Context, in *pb.Employee) (*pb.EmployeeId, error){
	str := fmt.Sprintf("CreateEmployee was invoked with %v\n", in)

	sng.DebugLogger.Debug(ctx, str, data)

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
	
	res, err:= empdb.MongoService.Collection.InsertOne(ctx, data)

	if err != nil{
		sng.ErrorLogger.Error(ctx,err,"Error while inserting data",data)

		return nil, status.Errorf(
		codes.Internal,
		fmt.Sprintf("Internal Error: %v\n", err),
		)
	}

	
	_, ok := res.InsertedID.(primitive.ObjectID)
	if !ok{
		sng.ErrorLogger.Error(ctx,err,"Cannot convert to OID ",data)
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