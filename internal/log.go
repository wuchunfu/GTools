package internal

import (
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
	once   sync.Once
)

// NewLogger ...
func NewLogger(logFile string) *logrus.Logger {
	once.Do(func() {
		logger = logrus.New()
		f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			panic("创建日志文件失败: " + err.Error())
		}
		writer := []io.Writer{f, os.Stdout}
		multiWriter := io.MultiWriter(writer...)
		logger.SetOutput(multiWriter)
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			DisableColors: false,
			FullTimestamp: true,
		})
	})
	return logger
}
