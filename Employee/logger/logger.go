package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


var DebugLogger SingletonLogger
var InfoLogger SingletonLogger
var WarnLogger SingletonLogger
var ErrorLogger SingletonLogger
var FatalLogger SingletonLogger

var data interface{}

var (
	logLevels = map[string]zapcore.Level{
		"DEBUG":	zap.DebugLevel,
		"INFO":		zap.InfoLevel,
		"WARN":		zap.WarnLevel,
		"ERROR":	zap.ErrorLevel,
		"FATAL":	zap.FatalLevel,
	}
)

type SingletonLogger struct{
	Logger *zap.Logger
	Service string
}

func NewEncoderConfig() zapcore.EncoderConfig{
	return zapcore.EncoderConfig{
		TimeKey: "timestamp",
		LevelKey: "category",
		NameKey: "logger",
		CallerKey: "caller",
		FunctionKey: zapcore.OmitKey,
		MessageKey: "msg",
		StacktraceKey: "stacktrace",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
}

func NewConfig(logLevel string) zap.Config{
	return zap.Config{
		Level: zap.NewAtomicLevelAt(logLevels[logLevel]),
		Encoding: "json",
		EncoderConfig: NewEncoderConfig(),
		OutputPaths: []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func InitSingletonLogger(ctx context.Context, service string){
	dlogger,_ := NewConfig("DEBUG").Build()
	DebugLogger = SingletonLogger{
			Logger: dlogger,
			Service: service,
	}

	ilogger,_ := NewConfig("INFO").Build()
	InfoLogger = SingletonLogger{
			Logger: ilogger,
			Service: service,
	}

	wlogger,_ := NewConfig("WARN").Build()
	WarnLogger = SingletonLogger{
			Logger: wlogger,
			Service: service,
	}

	elogger,_ := NewConfig("ERROR").Build()
	ErrorLogger = SingletonLogger{
			Logger: elogger,
			Service: service,
	}

	flogger,_ := NewConfig("FATAL").Build()
	FatalLogger = SingletonLogger{
			Logger: flogger,
			Service: service,
	}
}

// func (s *SingletonLogger) Log(level zapcore.Level, message string, fields ...zap.Field) {
// 	fields = append(fields, zap.String("service", SngService.Service))
//     SngService.Logger.Log(level, message, fields...)
// }

func (s *SingletonLogger) Debug(ctx context.Context, message string, data interface{}){
	DebugLogger.Logger.Debug(message,
		zap.String("service", DebugLogger.Service),
		zap.Any("data",data),
	)
}

func (s *SingletonLogger) Info(ctx context.Context, message string, data interface{}){
	InfoLogger.Logger.Info(message,
		zap.String("service", InfoLogger.Service),
		//zap.String("client", functions.GetValueFromContext(ctx, "client")),
		//zap.String("transactionID", functions.GetValueFromContext(ctx, "transactionID")),
		//zap.String("correlationID", functions.GetValueFromContext(ctx, "correlationID")),
		zap.Any("data",data),
	)
}

func (s *SingletonLogger) Warn(ctx context.Context, message string){
	WarnLogger.Logger.Warn(message,
		zap.String("service", WarnLogger.Service),
	)
}

func (s *SingletonLogger) Error(ctx context.Context, err error, errorMessage string, data interface{}){
	ErrorLogger.Logger.Error(errorMessage,
		zap.String("service", ErrorLogger.Service),
		//zap.String("errorCode", errorCode),
		zap.Error(err),
		zap.Any("data",data),
	)
}

func (s *SingletonLogger) Fatal(ctx context.Context, err error, errorMessage string, data interface{}){
	FatalLogger.Logger.Fatal(errorMessage,
		zap.String("service", FatalLogger.Service),
		//zap.String("errorCode", errorCode),
		zap.Error(err),
		zap.Any("data", data),
	)
}

// func (s *SingletonLogger) Sync(){
// 	SngService.Logger.Sync()
// }