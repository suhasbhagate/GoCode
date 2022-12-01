package main

import (
	"context"
	"log"

	pb "github.com/sbbhagate/GoCode/Employee/proto"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) PatchEmployee(ctx context.Context, in *pb.Employee) (*pb.UpdateEmployeeResponse, error) {
	log.Printf("PatchEmployee was invoked with %v\n", in)

	data := &Emp{}
	filter := bson.M{"empid": in.EmpId}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err !=nil{
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find Record with the Employee ID provided",
			)
	}

	ret := documentToEmployee(data)

	var fnm string
	if in.FirstName !=""{
		fnm = in.FirstName
	} else {
		fnm = ret.FirstName
	}

	var lnm string
	if in.LastName !=""{
		lnm = in.LastName
	} else {
		lnm = ret.LastName
	}

	var dsnm string
	if in.DisplayName !=""{
		dsnm = in.DisplayName
	} else {
		dsnm = ret.DisplayName
	}

	var des string
	if in.Designation !=""{
		des = in.Designation
	} else {
		des = ret.Designation
	}

	var dpt string
	if in.Department !=""{
		dpt = in.Department
	} else {
		dpt = ret.Department
	}

	var ag int32

	if in.Age > 0{
		ag = in.Age
	} else {
		ag = ret.Age
	}

	var sl float64
	if in.Salary > 0{
		sl = in.Salary
	} else {
		sl = ret.Salary
	}

	data1 := Emp{
		EmpId:       in.EmpId,
		FirstName:   fnm,
		LastName:    lnm,
		DisplayName: dsnm,
		Age:         ag,
		Salary:      sl,
		Designation: des,
		Department:  dpt,
	}

	res1, err1 := collection.UpdateOne(
		ctx,
		bson.M{"empid": in.EmpId},
		bson.M{"$set": data1},
	)

	if err1 != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}

	if res1.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find Employee with given Id",
		)
	}

	CreateMatView(ctx)

	return &pb.UpdateEmployeeResponse{Response: "Patched Successfully"}, nil
}
