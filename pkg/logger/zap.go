package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init(env string) *zap.Logger {
	var (
		err    error
		l      *zap.Logger
		config zap.Config
	)

	if env == "PROD" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l, err = config.Build()
	if err != nil {
		log.Fatalln("Error building logger", err)
		return nil
	}
	zap.ReplaceGlobals(l)
	l.Info("logger setup done")

	return l
}
