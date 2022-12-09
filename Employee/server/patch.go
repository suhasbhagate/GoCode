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

func (s *Server) PatchEmployee(ctx context.Context, in *pb.Employee) (*pb.UpdateEmployeeResponse, error) {
	str := fmt.Sprintf("PatchEmployee was invoked with %v\n", in)
	sng.DebugLogger.Debug(ctx, str, data)

	data := &Emp{}
	filter := bson.M{"empid": in.EmpId}

	res := sng.MongoService.Collection.FindOne(ctx, filter)

	if err := res.Decode(data); err !=nil{
		sng.ErrorLogger.Error(ctx,err,"Cannot find Record with the Employee ID provided",data)

		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find Record with the Employee ID provided",
			)
	}

	ret := documentToEmployee(data)

	var fnm string
	if in.FirstName !=""{
		fnm = in.FirstName
		sng.InfoLogger.Info(ctx, "First Name Updated", data)
	} else {
		fnm = ret.FirstName
	}

	var lnm string
	if in.LastName !=""{
		lnm = in.LastName
		sng.InfoLogger.Info(ctx, "Last Name Updated", data)
	} else {
		lnm = ret.LastName
	}

	var dsnm string
	if in.DisplayName !=""{
		dsnm = in.DisplayName
		sng.InfoLogger.Info(ctx, "Display Name Updated", data)
		
	} else {
		dsnm = ret.DisplayName
	}

	var des string
	if in.Designation !=""{
		des = in.Designation
		sng.InfoLogger.Info(ctx, "Designation Updated", data)
	} else {
		des = ret.Designation
	}

	var dpt string
	if in.Department !=""{
		dpt = in.Department
		sng.InfoLogger.Info(ctx, "Department Updated", data)
	} else {
		dpt = ret.Department
	}

	var ag int32

	if in.Age > 0{
		ag = in.Age
		sng.InfoLogger.Info(ctx, "Age Updated", data)
	} else {
		ag = ret.Age
	}

	var sl float64
	if in.Salary > 0{
		sl = in.Salary
		sng.InfoLogger.Info(ctx, "Salary Updated", data)
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

	res1, err1 := sng.MongoService.Collection.UpdateOne(
		ctx,
		bson.M{"empid": in.EmpId},
		bson.M{"$set": data1},
	)

	if err1 != nil {
		sng.ErrorLogger.Error(ctx,err1,"Could not update record",data)
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}

	if res1.MatchedCount == 0 {
		sng.ErrorLogger.Error(ctx,err1,"Cannot find Employee with given Id",data)
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find Employee with given Id",
		)
	}

	CreateMatView(ctx)

	return &pb.UpdateEmployeeResponse{Response: "Patched Successfully"}, nil
}