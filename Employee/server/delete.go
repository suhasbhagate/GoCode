package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	pb "github.com/sbbhagate/GoCode/Employee/proto"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteEmployee(ctx context.Context, in *pb.EmployeeId) (*pb.DeleteEmployeeResponse, error) {
	log.Printf("DeleteEmployee was invoked with %v\n", in)

	res, err := collection.DeleteOne(
		ctx,
		bson.M{"empid": in.Empid},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete Record in MongoDB %v\n", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find Record with given EmployeeId",
		)
	}

	res1, err := collection1.DeleteOne(
		ctx,
		bson.M{"empid": strconv.Itoa(int(in.Empid))},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete Record in MongoDB %v\n", err),
		)
	}

	if res1.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find Record with given EmployeeId",
		)
	}

	return &pb.DeleteEmployeeResponse{Response: "Deleted Successfully"}, nil
}
