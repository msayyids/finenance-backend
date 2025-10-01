package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logger() *zap.Logger {

	log := zap.NewDevelopmentConfig()

	//if env == "production" {
	//	Logger, err = zap.NewProduction() // log JSON, high performance
	//}

	log.DisableStacktrace = true
	log.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	log.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	log.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	logger, err := log.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
