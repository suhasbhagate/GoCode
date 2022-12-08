package main
 
import "go.uber.org/zap"
 
func main() {
   logger, _ := zap.NewProduction()
 
   logger.Info("INFO log level message")
   logger.Warn("Warn log level message")
   logger.Error("Error log level message")
  }