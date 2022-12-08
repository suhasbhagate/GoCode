package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load("local.env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    fmt.Println(os.Getenv("GRPCADDR"))
    fmt.Println(os.Getenv("HTTPADDR"))

    // os.Setenv("FOO", "1")
    // fmt.Println("FOO:", os.Getenv("FOO"))
    // fmt.Println("BAR:", os.Getenv("BAR"))
}