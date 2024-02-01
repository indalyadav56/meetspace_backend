package services

import (
	"os"

	"github.com/sirupsen/logrus"
)

// LogService provides a log interface
type LogService interface {
    Debugf(format string, args ...interface{})
    Infof(format string, args ...interface{}) 
    Errorf(format string, args ...interface{})
}

type logService struct {
    log *logrus.Logger
}

// NewLogService creates a log service
func NewLogService() LogService {
    l := logrus.New()
    l.SetOutput(os.Stdout)
    l.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })

    return &logService{log: l}
}

// Debugf logs debug msg 
func (l *logService) Debugf(format string, args ...interface{}) {
    l.log.Debugf(format, args...)
}

// Infof logs info msg
func (l *logService) Infof(format string, args ...interface{}) {
    l.log.Infof(format, args...) 
}

// Errorf logs error msg
func (l *logService) Errorf(format string, args ...interface{}) {
    l.log.Errorf(format, args...)
}