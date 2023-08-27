package main

import (
	"log"
	"os"

	"go.uber.org/zap"
)

func main() {
	file, err := os.OpenFile("D:\\GoCode\\GoCodeAugust2\\LoggingDemo\\log1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)
	log.Printf("Set Log Output File")

	conf := zap.NewProductionConfig()
	conf.OutputPaths = append(conf.OutputPaths, "D:\\GoCode\\GoCodeAugust2\\LoggingDemo\\log2.txt")
	logger, err := conf.Build()
	if err != nil {
		log.Println(err)
	}
	logger.Info("Set Log Output File")
}
