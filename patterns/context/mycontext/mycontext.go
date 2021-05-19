package mycontext

import (
	"context"
	"net/http"
)

type key string

const (
	someValueKey        = key("some_value")
	someValueRequestKey = "X-Some-Value"
)

func WithSomeValue(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, someValueKey, value)
}

func SomeValue(ctx context.Context) string {
	someValue, ok := ctx.Value(someValueKey).(string)
	if !ok {
		return ""
	}
	return someValue
}

func WithSomeValueRequest(req *http.Request) *http.Request {
	req.Header.Set(someValueRequestKey, SomeValue(req.Context()))
	return req
}

func SomeValueFromRequest(req *http.Request) string {
	return req.Header.Get(someValueRequestKey)
}
