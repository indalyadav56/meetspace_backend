package services

import (
	"os"

	"github.com/sirupsen/logrus"
)

// LogService provides a log interface
type LoggerService struct {
    log *logrus.Logger
}

// NewLogService creates a log service
func NewLoggerService() *LoggerService {
    l := logrus.New()
    l.SetOutput(os.Stdout)
    l.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })

    return &LoggerService{log: l}
}

// Debugf logs debug msg 
func (l *LoggerService) Debugf(format string, args ...interface{}) {
    l.log.Debugf(format, args...)
}

// Infof logs info msg
func (l *LoggerService) Infof(format string, args ...interface{}) {
    l.log.Infof(format, args...) 
}

// Errorf logs error msg
func (l *LoggerService) Errorf(format string, args ...interface{}) {
    l.log.Errorf(format, args...)
}