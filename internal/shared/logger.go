package shared

import (
	"go.uber.org/zap"
	"sync"
)

var once sync.Once

var (
	logger *zap.SugaredLogger
)

func GetSugarLogger() *zap.SugaredLogger {
	var zapLogger *zap.Logger

	once.Do(func() {
		zapLogger, _ = zap.NewProduction()
		logger = zapLogger.Sugar()
		defer zapLogger.Sync()
	})

	return logger
}
