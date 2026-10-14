package logger

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// WriterConfig holds configuration for log rotation based on environments
type WriterConfig struct {
	Filename   string
	MaxSize    int  // megabytes
	MaxBackups int  // number of old backups to keep
	MaxAge     int  // days
	Compress   bool // whether to compress disabled logs
}

// getLogWriter builds a write syncer utilizing lumberjack for log rotation
func getLogWriter(cfg WriterConfig) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
