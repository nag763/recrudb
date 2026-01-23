package logger

import (
	"nag763/recrudb/api/internal/pkg/utils"

	"go.uber.org/zap"
)

func NewLoggerService() *zap.SugaredLogger {
	var logger *zap.Logger
	if utils.GetEnvOrDefault("ENV", "PROD") == "DEV" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}
	return logger.Sugar()
}
