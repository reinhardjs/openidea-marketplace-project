package config

import (
	"github.com/openidea-marketplace/domain"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type logrusLogger struct {
	logger *logrus.Logger
}

func NewLogger(viper *viper.Viper) domain.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.Level(viper.GetInt32("log.level")))
	logger.SetFormatter(&logrus.JSONFormatter{})

	return &logrusLogger{logger: logger}
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *logrusLogger) Printf(message string, args ...interface{}) {
	l.logger.Tracef(message, args...)
}
