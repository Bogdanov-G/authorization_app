// logger package provides:
// configuration initialization and configuration for logging tool;
// methods (wrappers for logger tool) which cover the most frequent
// error’s cases.
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// init() constructs logger with specified settings.
func init() {
	var err error

	config := zap.NewProductionConfig()
	enccoderConfig := zap.NewProductionEncoderConfig()
	enccoderConfig.TimeKey = "timestamp"
	enccoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	enccoderConfig.StacktraceKey = "" // disable stacktrace info showing
	config.EncoderConfig = enccoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

// Info() makes InfoLevel log records with a given message and key-value
// pairs given in fields.
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Debug() makes DebugLevel log records with a given message and key-value
// pairs given in fields.
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// Error() makes ErrorLevel log records with a given message and key-value
// pairs given in fields.
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

// Fatal() makes FatalLevel log records with a given message and key-value
// pairs given in fields.
func Fatal(message string, fields ...zap.Field) {
	log.Fatal(message, fields...)
}
