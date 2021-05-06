package mycontext

import "context"

type key string

const (
	someValueKey = key("some_value")
)

func New(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, someValueKey, value)
}

func SomeValue(ctx context.Context) string {
	someValue, ok := ctx.Value(someValueKey).(string)
	if !ok {
		return ""
	}
	return someValue
}
