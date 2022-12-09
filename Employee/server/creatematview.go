package main

import (
	"context"
	sng "github.com/sbbhagate/GoCode/Employee/singleton"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateMatView(ctx context.Context) {
	pipeln, err := sng.MongoService.Collection.Aggregate(ctx,
		bson.A{
			bson.D{
				{Key: "$project",
					Value: bson.D{
						{Key: "empid", Value: bson.D{{Key: "$toString", Value: "$empid"}}},
						{Key: "firstname", Value: "$firstname"},
						{Key: "lastname", Value: "$lastname"},
						{Key: "displayname", Value: "$displayname"},
						{Key: "age", Value: bson.D{{Key: "$toString", Value: "$age"}}},
						{Key: "salary", Value: bson.D{{Key: "$toString", Value: "$salary"}}},
						{Key: "designation", Value: "$designation"},
						{Key: "department", Value: "$department"},
					},
				},
			},
			bson.D{
				{Key: "$merge",
					Value: bson.D{
						{Key: "into", Value: "emp_mat_view"},
						{Key: "whenMatched", Value: "replace"},
					},
				},
			},
		})
	if err != nil {
		sng.FatalLogger.Fatal(ctx,err,"Error while creating Materialized View",data)
	}

	defer pipeln.Close(ctx)
}