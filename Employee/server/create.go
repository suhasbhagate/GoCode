package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/sbbhagate/GoCode/Employee/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateEmployee(ctx context.Context, in *pb.Employee) (*pb.EmployeeId, error){
	log.Printf("CreateEmployee was invoked with %v\n", in)

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
	
	res, err:= collection.InsertOne(ctx, data)
	if err != nil{
		return nil, status.Errorf(
		codes.Internal,
		fmt.Sprintf("Internal Error: %v\n", err),
		)
	}

	
	_, ok := res.InsertedID.(primitive.ObjectID)
	if !ok{
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