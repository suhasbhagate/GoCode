package singleton

import (
	"context"
	"fmt"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SngService SingletonLogger
var MongoService MongoDB

type SingletonLogger struct{
	Logger *zap.Logger
}

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
		SngService.Fatal("Error while creating MongoDB Client %s",zap.Error(err))
	} else{
		SngService.Info("MongoDB Client created successfully")
	}

	err = Mongoclient.Connect(context.Background())
	if err != nil {
		SngService.Fatal("Error while connecting to MongoDB %s",zap.Error(err))
	} else{
		SngService.Info("Connection to MongoDB is successful")
	}
	
	MongoService = MongoDB{
		Client:	Mongoclient,
		Collection: Mongoclient.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_COLLECTION")),
		MatView: Mongoclient.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_MATVIEW")),
	}
}

func (s *MongoDB) DisconnectMongoDB(ctx context.Context){
	s.Client.Disconnect(ctx)
}

func (s *MongoDB) PingMongoClient() error{
	return s.Client.Ping(context.TODO(),readpref.Primary())
}

func InitSingletonLogger(ctx context.Context){
	filename := "logs.log"
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	SngService = SingletonLogger{
		Logger: zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)),
	}
}

func (s *SingletonLogger) Info(message string, fields ...zap.Field){
	SngService.Logger.Info(message, fields...)
}

func (s *SingletonLogger) Warn(message string, fields ...zap.Field){
	SngService.Logger.Warn(message, fields...)
}

func (s *SingletonLogger) Error(message string, fields ...zap.Field){
	SngService.Logger.Error(message, fields...)
}

func (s *SingletonLogger) Fatal(message string, fields ...zap.Field){
	SngService.Logger.Fatal(message, fields...)
}

func (s *SingletonLogger) Debug(message string, fields ...zap.Field){
	SngService.Logger.Debug(message, fields...)
}

func (s *SingletonLogger) Sync(){
	SngService.Logger.Sync()
}