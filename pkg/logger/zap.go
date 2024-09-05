package logger

import (
	"go.uber.org/zap"
)

var l *zap.Logger

func Init(env string) {
	var err error
	if env == "PROD" {
		l, err = zap.NewProduction()
	} else {
		l, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(l)
	l.Info("logger setup done")
}

func Info(msg string, fields ...zap.Field) {
	l.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	l.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	l.Debug(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	l.Fatal(msg, fields...)
}

func Sync() error {
	return l.Sync()
}
