package config

import (
	"go.uber.org/zap"
)

func NewLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}

func CloseLogger(logger *zap.SugaredLogger) {
	_ = logger.Sync()
}
