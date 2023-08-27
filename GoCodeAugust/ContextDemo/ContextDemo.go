package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Name", "Saksham")
	fmt.Println(ctx.Value("Name"))
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
}
