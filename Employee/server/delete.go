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

func (s *Server) DeleteEmployee(ctx context.Context, in *pb.EmployeeId) (*pb.DeleteEmployeeResponse, error) {
	str := fmt.Sprintf("DeleteEmployee was invoked with %v\n", in)
	sng.DebugLogger.Debug(ctx, str, data)

	res, err := sng.MongoService.Collection.DeleteOne(
		ctx,
		bson.M{"empid": in.Empid},
	)

	if err != nil {
		sng.ErrorLogger.Error(ctx,err,"Cannot Delete Record in MongoDB",data)

		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot Delete Record in MongoDB %v\n", err),
		)
	}

	if res.DeletedCount == 0 {
		sng.ErrorLogger.Error(ctx,err,"Cannot Find Record with given EmployeeId",data)

		return nil, status.Errorf(
			codes.NotFound,
			"Cannot Find Record with given EmployeeId",
		)
	}

	res1, err := sng.MongoService.MatView.DeleteOne(
		ctx,
		bson.M{"empid": strconv.Itoa(int(in.Empid))},
	)

	if err != nil {
		sng.ErrorLogger.Error(ctx,err,"Cannot Delete Record in MongoDB",data)

		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot Delete Record in MongoDB %v\n", err),
		)
	}

	if res1.DeletedCount == 0 {
		sng.ErrorLogger.Error(ctx,err,"Cannot Find Record with given EmployeeId",data)

		return nil, status.Errorf(
			codes.NotFound,
			"Cannot Find Record with given EmployeeId",
		)
	}

	return &pb.DeleteEmployeeResponse{Response: "Record Deleted Successfully"}, nil
}