package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewProductionLogger builds and returns a Logger designed for production environment.
func NewProductionLogger() Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(cfg)

	writeSyncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(getLogWriter(WriterConfig{
			Filename:   "logs/app-production.log",
			MaxSize:    50,
			MaxBackups: 30,
			MaxAge:     30,
			Compress:   true,
		})),
	)

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &zapLogger{
		logger: logger,
	}
}
