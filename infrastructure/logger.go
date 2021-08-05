package infrastructure

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Zap *zap.SugaredLogger
}

func NewLogger() Logger {
	config := zap.NewDevelopmentConfig()

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, _ := config.Build()

	sugar := logger.Sugar()

	return Logger{
		Zap: sugar,
	}
}
