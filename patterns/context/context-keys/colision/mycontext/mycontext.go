package mycontext

import (
	"context"
)

const ctxKey = "ctxKey"

func WithSomeValue(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, ctxKey, value)
}

func SomeValue(ctx context.Context) string {
	someValue, ok := ctx.Value(ctxKey).(string)
	if !ok {
		return ""
	}
	return someValue
}
