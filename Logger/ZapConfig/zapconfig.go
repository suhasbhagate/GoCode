package main
 
import (
   "encoding/json" 
   "go.uber.org/zap"
)
 
func main() {
   sampleJSON := []byte(`{
       "level" : "info",
       "encoding": "json",
       "outputPaths":["stdout", "log.log"],
       "errorOutputPaths":["stderr"],
       "encoderConfig": {
           "messageKey":"message",
           "levelKey":"level",
           "levelEncoder":"lowercase"
       }
   }`)
 
   var cfg zap.Config
 
   if err := json.Unmarshal(sampleJSON, &cfg); err != nil {
       panic(err)
   }
 
   logger, err := cfg.Build()
 
   if err != nil {
       panic(err)
   }
   defer logger.Sync()
 
 
   logger.Info("INFO log level message")
   logger.Warn("Warn log level message")
   logger.Error("Error log level message")
}