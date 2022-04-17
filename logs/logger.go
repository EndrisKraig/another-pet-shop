package logs

import (
	"os"

	"go.uber.org/zap"
)

const (
	logPath = "./logs/go.log"
)

var Logger *zap.Logger

func InitLogger() {
	os.OpenFile(logPath, os.O_RDONLY|os.O_CREATE, 0666)
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", logPath}
	l, err := c.Build()
	if err != nil {
		panic(err)
	}
	Logger = l
	Logger.Info("Logger successfully initialized")
}
