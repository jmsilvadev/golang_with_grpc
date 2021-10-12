package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log zap.Logger object
var (
	Log *zap.Logger
)

// SetLogger defines the zap logger to used by dependency injection and overrides the standard log
func SetLogger() *zap.Logger {

	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zapcore.ErrorLevel
	})

	errorFatalLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		nanos := t.UnixNano()
		millis := nanos / int64(time.Millisecond)
		enc.AppendInt64(millis)
	}

	consoleEncoder := zapcore.NewJSONEncoder(config)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, errorFatalLevel),
		zapcore.NewCore(consoleEncoder, consoleDebugging, infoLevel),
	)

	Log = zap.New(core)
	zap.RedirectStdLog(Log)
	return Log
}
