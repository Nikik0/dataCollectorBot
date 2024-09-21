package logger

import (
	"go.uber.org/zap"
	"log"
)

var logger *zap.Logger

func init() {
	newLogger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to init zap logger")
	}
	logger = newLogger
}

func Fatal(msg string, additionalData ...interface{}) {
	sugar := logger.Sugar()
	sugar.Fatal(msg, additionalData)
}

func Error(msg string, additionalData ...interface{}) {
	sugar := logger.Sugar()
	sugar.Error(msg, additionalData)
}

func Warn(msg string, additionalData ...interface{}) {
	sugar := logger.Sugar()
	sugar.Warn(msg, additionalData)
}

func Info(msg string, additionalData ...interface{}) {
	sugar := logger.Sugar()
	sugar.Info(msg, additionalData)
}

func Debug(msg string, additionalData ...interface{}) {
	sugar := logger.Sugar()
	sugar.Debug(msg, additionalData)
}
