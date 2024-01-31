package main

import (
	"github.com/lni/dragonboat/v3/logger"
	"go.uber.org/zap"
)

type LogLevel = logger.LogLevel

type Mylogger struct {
	*zap.Logger
}

func (l *Mylogger) SetLevel(level LogLevel) {
}

func (l *Mylogger) Warningf(format string, args ...interface{}) {
	l.Logger.Sugar().Warnf(format, args...)
}

func (l *Mylogger) Debugf(format string, args ...interface{}) {
	l.Logger.Sugar().Debugf(format, args...)
}

func (l *Mylogger) Errorf(format string, args ...interface{}) {
	l.Logger.Sugar().Errorf(format, args...)
}

func (l *Mylogger) Infof(format string, args ...interface{}) {
	l.Logger.Sugar().Infof(format, args...)
}

func (l *Mylogger) Panicf(format string, args ...interface{}) {
	l.Logger.Sugar().Panicf(format, args...)
}

var _ logger.ILogger = (*Mylogger)(nil)

var factory = func(pkgName string) logger.ILogger {
	logger, _ := zap.NewProduction()
	return &Mylogger{
		Logger: logger,
	}
}

func main() {
	logger.SetLoggerFactory(logger.Factory(factory))
	logger.GetLogger("raft").SetLevel(logger.ERROR)

}
