package main
import (
"log"
"net/rpc"
) 
type Args struct {
A, B int
} 
func main() {
serverAddress := "127.0.0.1"
client, err := rpc.DialHTTP("tcp", serverAddress+":4422")
if err != nil {
log.Fatal("Dial failed, err:", err)
} 
args := Args{3, 4}
var reply int
client.Call("MuliplyService.Do", args, &reply)
log.Printf(" %d*%d=%d", args.A, args.B, reply)
}