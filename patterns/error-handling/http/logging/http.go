package logging

import (
	"log"

	"go.uber.org/zap"
)

type writerFunc func(p []byte) (n int, err error)

func (w writerFunc) Write(p []byte) (n int, err error) {
	return w(p)
}

// HTTPServerLogger returns a log.Logger that logs HTTP server internal errors
func HTTPServerLogger() *log.Logger {
	logger := Logger().With(zap.String("source", "http_server"))
	return log.New(
		writerFunc(func(p []byte) (n int, err error) {
			logger.Error(string(p))
			return len(p), nil
		}),
		"",
		0,
	)
}
