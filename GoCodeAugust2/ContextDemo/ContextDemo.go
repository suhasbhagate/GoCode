package main

import(
	"fmt"
	"context"
)

func main(){
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Name","Saksham")
	fmt.Println(ctx.Value("Name"))
	ctx,cancel := context.WithCancel(ctx)
	defer cancel()
}