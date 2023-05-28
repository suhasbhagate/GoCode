package db

import (
	"context"
	"fmt"
	"os"
	sng "github.com/suhasbhagate/GoCode/Employee/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoService MongoDB
var data interface{}

type MongoDB struct{
	Client		*mongo.Client
	Collection	*mongo.Collection
	MatView		*mongo.Collection
}

func InitMongoDB(ctx context.Context){
	mongoString := os.Getenv("MONGO_URL")
	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoURI := fmt.Sprintf(mongoString, mongoUsername, mongoPassword)

	Mongoclient, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		sng.FatalLogger.Fatal(ctx,err,"Error while creating MongoDB Client",data)
	} else{
		sng.InfoLogger.Info(ctx, "MongoDB Client created successfully", data)
	}

	err = Mongoclient.Connect(context.Background())
	if err != nil {
		sng.FatalLogger.Fatal(ctx,err,"Error while connecting to MongoDB",data)
	} else{
		sng.InfoLogger.Info(ctx, "Connection to MongoDB is successful",data)
	}
	
	MongoService = MongoDB{
		Client:	Mongoclient,
		Collection: Mongoclient.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_COLLECTION")),
		MatView: Mongoclient.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_MATVIEW")),
	}
	//fmt.Println(MongoService.Collection)
	//fmt.Println(MongoService.MatView)	
}

func (s *MongoDB) DisconnectMongoDB(ctx context.Context){
	s.Client.Disconnect(ctx)
}

func (s *MongoDB) PingMongoClient() error{
	return s.Client.Ping(context.TODO(),readpref.Primary())
}