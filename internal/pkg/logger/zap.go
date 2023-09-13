package logger

import "go.uber.org/zap"

var sugar *zap.SugaredLogger

func Zap() *zap.SugaredLogger {
	return sugar
}

func InitZap() {
	sugar = zap.NewExample().Sugar()
}
