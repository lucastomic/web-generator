package logging

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var relativeLogFilePath = "../logs/server.log"

type LogrusLogger struct {
	loggers []*logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	file, err := os.OpenFile(relativeLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Error opening log file", err)
	}
	fileLogger := logrus.New()
	fileLogger.Out = file

	ttyLogger := logrus.New()
	return &LogrusLogger{
		loggers: []*logrus.Logger{ttyLogger, fileLogger},
	}
}

func (l *LogrusLogger) Request(ctx context.Context, req *http.Request, status int) {
	msg := fmt.Sprintf(
		"%s | %s | %d",
		time.Now().Format("2006-01-02 15:04:05"),
		req.UserAgent(),
		status,
	)

	for _, logger := range l.loggers {
		logger.WithField("requestID", ctx.Value("requestID")).Info(msg)
	}
}

func (l *LogrusLogger) Info(args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Info(args...)
	}
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Warn(args...)
	}
}

func (l *LogrusLogger) Error(args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Error(args...)
	}
}

func (l *LogrusLogger) Fatal(args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Fatal(args...)
	}
}
