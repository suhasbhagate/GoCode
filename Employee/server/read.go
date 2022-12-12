package main

import (
	"context"
	"fmt"
	"strconv"
	pb "github.com/sbbhagate/GoCode/Employee/proto"
	sng "github.com/sbbhagate/GoCode/Employee/logger"
	empdb "github.com/sbbhagate/GoCode/Employee/db"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *Server) ReadEmployee(ctx context.Context, in *pb.Employee) (*pb.Employees, error) {
	str := fmt.Sprintf("ReadEmployee is invoked with %v\n", in)
	sng.DebugLogger.Debug(ctx, str, data)

	CreateMatView(ctx)

	var ppln bson.A
	var docFilters bson.A

	if in.EmpId<=0 && in.FirstName=="" && in.LastName=="" && in.DisplayName=="" && in.Department=="" && in.Designation=="" && in.Salary<=0 && in.Age<=0{
		sng.WarnLogger.Warn(ctx, "No Query Parameter is Entered, Please Enter atleast one Query Parameter")
	} else {
	if in.EmpId > 0{
		docFilters = append(docFilters, bson.D{{Key: "text",Value: bson.D{{Key: "query", Value: strconv.Itoa(int(in.EmpId))},{Key: "path",Value: "empid"}}}})
	}
	if in.FirstName !=""{
		docFilters = append(docFilters, bson.D{{Key: "text",Value: bson.D{{Key: "query", Value: in.FirstName},{Key: "path",Value: "firstname"}}}})
	}
	if in.LastName !=""{
		docFilters = append(docFilters, bson.D{{Key: "text",Value: bson.D{{Key: "query", Value: in.LastName},{Key: "path",Value: "lastname"}}}})
	}
	if in.DisplayName !=""{
		docFilters = append(docFilters, bson.D{{Key: "text",Value: bson.D{{Key: "query", Value: in.DisplayName},{Key: "path",Value: "displayname"}}}})
	}
	if in.Age > 0{
		docFilters = append(docFilters, bson.D{{Key: "text",Value: bson.D{{Key: "query", Value: strconv.Itoa(int(in.Age))},{Key: "path",Value: "age"}}}})
	}
	if in.Salary > 0{
		docFilters = append(docFilters, bson.D{{Key: "text",Value: bson.D{{Key: "query", Value: strconv.Itoa(int(in.Salary))},{Key: "path",Value: "salary"}}}})
	}
	if in.Designation !=""{
		docFilters = append(docFilters, bson.D{{Key: "text",Value: bson.D{{Key: "query", Value: in.Designation},{Key: "path",Value: "designation"}}}})
	}
	if in.Department !=""{
		docFilters = append(docFilters, bson.D{{Key: "text",Value: bson.D{{Key: "query", Value: in.Department},{Key: "path",Value: "department"}}}})
	}

	compound := bson.D{{Key: "filter",Value: docFilters}}

	ppln = append(ppln, bson.D{{Key: "$search",Value: bson.D{{Key: "index",Value: "EmployeeIndex"},{Key: "compound",Value: compound}}}})
	
	pipeline, err := empdb.MongoService.MatView.Aggregate(ctx,ppln)
	if err != nil {
		sng.FatalLogger.Fatal(ctx,err,"Error while searching data",data)
	}

	defer pipeline.Close(ctx)

	emplist := &pb.Employees{}

	for pipeline.Next(ctx) {
		data := &Emp1{}
		err := pipeline.Decode(data)

		if err != nil {
			sng.ErrorLogger.Error(ctx,err,"Error while retrieving data",data)
			return nil, status.Errorf(
				codes.NotFound,
				err.Error(),
			)
		}
		emplist.Emps = append(emplist.Emps, documentToEmployee1(data))
	}
	if err := pipeline.Err(); err != nil {
		sng.ErrorLogger.Error(ctx,err,"Error while retrieving data",data)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return emplist, nil
	}
	return nil,nil
}