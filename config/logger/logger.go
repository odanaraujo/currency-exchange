package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var (
	log         *zap.Logger
	LOG_OUTPUT  = "LOG_OUTPUT"
	LOG_LEVEL   = "LOG_LEVEL"
	LEVEL_KEY   = "LEVEL"
	TIME_KEY    = "time"
	MESSAGE_KEY = "message"
	ENCODING    = "json"
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutPutLogs()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    ENCODING,
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     LEVEL_KEY,
			TimeKey:      TIME_KEY,
			MessageKey:   MESSAGE_KEY,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, _ = logConfig.Build() // builda o log
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync()
}

func getOutPutLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
