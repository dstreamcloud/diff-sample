package uberzap

import (
	"github.com/dstreamcloud/diff-sample/pkg/logger"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func New() logger.Logger {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return &Logger{
		logger: l,
	}
}

func (l *Logger) Log(msg string) error {
	l.logger.Info(msg)
	return nil
}
