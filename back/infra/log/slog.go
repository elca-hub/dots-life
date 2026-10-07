package log

import (
	"back/adapter/logger"
	"back/mode"
	"log/slog"
	"os"
)

type slogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger(m string) (logger.Logger, error) {
	logLevel := new(slog.LevelVar)

	switch m {
	case mode.DevelopmentMode, mode.StagingMode:
		logLevel.Set(slog.LevelDebug)
	case mode.ProductionMode:
		logLevel.Set(slog.LevelInfo)
	}
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))

	return &slogLogger{logger: log}, nil
}

func (l *slogLogger) Debugf(format string, args ...any) {
	l.logger.Debug(format, args...)
}

func (l *slogLogger) Infof(format string, args ...any) {
	l.logger.Info(format, args...)
}

func (l *slogLogger) Warnf(format string, args ...any) {
	l.logger.Warn(format, args...)
}

func (l *slogLogger) Errorf(format string, args ...any) {
	l.logger.Error(format, args...)
}

func (l *slogLogger) WithFields(keyValues logger.Fields) logger.Logger {
	for k, v := range keyValues {
		l.logger = l.logger.With(k, v)
	}

	return &slogLogger{logger: l.logger}
}

func (l *slogLogger) WithError(err error) logger.Logger {
	log := l.logger.With("error", err.Error())

	return &slogLogger{logger: log}
}
