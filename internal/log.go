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
		// 分别输出到控制台和log日志
		multiWriter := io.MultiWriter(writer...)
		// log日志样式格式化
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			DisableColors: false,
			TimestampFormat:"2006-01-02 15:04:05",  //时间格式
			FullTimestamp: true,
		})
		logger.SetOutput(multiWriter)
	})
	return logger
}
