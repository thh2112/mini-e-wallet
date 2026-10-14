package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewDevelopmentLogger() Logger {
	consoleCfg := zap.NewDevelopmentEncoderConfig()
	consoleCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleCfg)

	fileCfg := zap.NewDevelopmentEncoderConfig()
	fileCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	fileEncoder := zapcore.NewConsoleEncoder(fileCfg)

	fileSyncer := zapcore.AddSync(getLogWriter(WriterConfig{
		Filename:   "logs/app-development.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     7,
		Compress:   false,
	}))

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(fileEncoder, fileSyncer, zapcore.DebugLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &zapLogger{
		logger: logger,
	}
}
