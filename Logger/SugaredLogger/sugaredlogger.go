package main

import (
	"log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("Jan 02 15:04:05.000000000")
	config.EncoderConfig.StacktraceKey = "" // to hide stacktrace info

	logger, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}

	log := logger.Sugar()

	log.Info("INFO log level message")
	log.Warn("Warn log level message")
	log.Error("Error log level message")
}