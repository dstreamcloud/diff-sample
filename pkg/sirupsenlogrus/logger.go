package sirupsenlogrus

import (
	"github.com/dstreamcloud/diff-sample/pkg/logger"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func New() logger.Logger {
	return &Logger{
		logger: logrus.New(),
	}
}

func (l *Logger) Log(msg string) error {
	l.logger.Println(msg)
	return nil
}
