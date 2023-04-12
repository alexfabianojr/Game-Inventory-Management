package logger

import "go.uber.org/zap"

func Get() *zap.SugaredLogger {
	log, _ := zap.NewProduction()
	return log.Sugar()
}
