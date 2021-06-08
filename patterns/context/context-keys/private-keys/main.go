package main

import (
	"context"
	"fmt"

	"concurrency/patterns/context/context-keys/private-keys/mycontext"
)

const ctxKey = "ctxKey"

func main() {
	ctx := mycontext.WithSomeValue(context.Background(), "main")

	ctx = req(ctx)
	fmt.Println("ctx main value:", mycontext.SomeValue(ctx))

	val := ctx.Value(ctxKey).(string)
	fmt.Println("ctx req value:", val)
}

func req(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey, "req")
}
