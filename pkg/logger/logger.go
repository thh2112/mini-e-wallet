package logger

import (
	"context"
	"strings"

	"go.uber.org/zap"
)

type Logger interface {
	WithFields(fields map[string]any) Logger
	Debug(ctx context.Context, message string)
	Info(ctx context.Context, message string)
	Warn(ctx context.Context, message string)
	Error(ctx context.Context, message string)
	Fatal(ctx context.Context, message string)
	Panic(ctx context.Context, message string)
}

type zapLogger struct {
	logger *zap.Logger
}

func (l *zapLogger) WithFields(fields map[string]any) Logger {
	return &zapLogger{
		logger: l.logger.With(zap.Any("fields", fields)),
	}
}

func (l *zapLogger) Debug(ctx context.Context, message string) {
	l.logger.Debug(message)
}

func (l *zapLogger) Info(ctx context.Context, message string) {
	l.logger.Info(message)
}

func (l *zapLogger) Warn(ctx context.Context, message string) {
	l.logger.Warn(message)
}

func (l *zapLogger) Error(ctx context.Context, message string) {
	l.logger.Error(message)
}

func (l *zapLogger) Fatal(ctx context.Context, message string) {
	l.logger.Fatal(message)
}

func (l *zapLogger) Panic(ctx context.Context, message string) {
	l.logger.Panic(message)
}

// GetLoggerInstance acts as a simple factory method to retrieve the correct logger implementation based on the environment mode.
func GetLoggerInstance(mode string) Logger {
	switch strings.ToLower(mode) {
	case "production":
		return NewProductionLogger()
	case "development":
		return NewDevelopmentLogger()
	default:
		return NewDevelopmentLogger()
	}
}

func InitLogger() Logger {
	return GetLoggerInstance("development") // Default to development
}
