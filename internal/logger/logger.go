package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var slogger *zap.SugaredLogger

func init() {
	if slogger == nil {
		NewLogger(zap.InfoLevel)
	}
}

func NewLogger(level zapcore.Level, opts ...zap.Option) {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Encoding:    "json",
		Development: false,
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "ts",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}

	logger, err := cfg.Build(opts...)
	if err != nil {
		log.Fatalf("cannot init logger: %v", err)
	}

	slogger = logger.Sugar()
	return
}

func Debug(args ...interface{}) {
	slogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	slogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	slogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	slogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	slogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	slogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	slogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	slogger.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	slogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	slogger.Fatalf(template, args...)
}

func FatalKV(msg string, kvPairs ...interface{}) {
	slogger.Fatalw(msg, kvPairs...)
}
