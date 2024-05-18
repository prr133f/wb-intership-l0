package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

func NewZap() *zap.Logger {
	logger := zap.Must(zap.NewDevelopment())
	if strings.EqualFold(os.Getenv("APP_STATUS"), "release") {
		logger = zap.Must(zap.NewProduction())
	}
	return logger
}
