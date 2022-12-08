package main
 
import (
   "log"
   "os"
)
 
func main() {
   // Log to the console
   log.Println("Message to log in the console!")
   // Log in a log file
   file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
   if err != nil {
       log.Fatal(err)
   }
   defer file.Close()
   log.SetOutput(file)
   log.Println("Message to log in a log file!")
}