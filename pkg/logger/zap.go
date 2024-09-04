package logger

import (
	"go.uber.org/zap"
)

var L *zap.Logger

func Init(env string) {
	var err error
	if env == "PROD" {
		L, err = zap.NewProduction()
	} else {
		L, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(L)
	L.Info("Logger setup done")
}
