package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var l *zap.Logger

func Init() error {
	var logLevel zapcore.Level
	err := logLevel.Set("debug")
	if err != nil {
		return err
	}

	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(logLevel)
	zapConfig.OutputPaths = []string{"stdout", "app.log"}
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := zapConfig.Build()
	defer func() {
		_ = logger.Sync()
	}()
	l = logger
	return err
}

func Logger() *zap.Logger {
	return l
}
