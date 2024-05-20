package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

// NewZap returns a new instance of the zap.Logger.
//
// It creates a new zap.Logger using the zap.NewDevelopment() function.
// If the environment variable "APP_STATUS" is set to "release",
// it creates a new zap.Logger using the zap.NewProduction() function.
func NewZap() *zap.Logger {
	logger := zap.Must(zap.NewDevelopment())
	if strings.EqualFold(os.Getenv("APP_STATUS"), "release") {
		logger = zap.Must(zap.NewProduction())
	}
	return logger
}
