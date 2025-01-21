package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"vega-server/pkg/config"

	"os"
	"time"
)

type Logger struct {
	// Embedding
	*zap.Logger
}

func NewLogger(config *config.Config) *Logger {
	// Initialize zapcore.Encoder
	var encoder zapcore.Encoder
	if config.GetString("log.format") == "console" {
		encoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "timestamp",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     customTimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
	} else {
		encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "timestamp",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     customTimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
	}
	// Initialize zapcore.WriteSyncer
	lumberjackLogger := &lumberjack.Logger{
		Filename:   config.GetString("log.file_name"),
		MaxSize:    config.GetInt("log.max_size"),
		MaxBackups: config.GetInt("log.max_backups"),
		MaxAge:     config.GetInt("log.max_age"),
		Compress:   config.GetBool("log.compress"),
	}
	// Initialize zapcore.LevelEnabler
	var levelEnabler zapcore.LevelEnabler
	switch level := config.GetString("log.level"); level {
	case "Debug":
		levelEnabler = zap.DebugLevel
	case "Info":
		levelEnabler = zap.InfoLevel
	case "Warn":
		levelEnabler = zap.WarnLevel
	case "Error":
		levelEnabler = zap.ErrorLevel
	default:
		levelEnabler = zap.DebugLevel
	}
	// Initialize zapcore.Core
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberjackLogger)),
		levelEnabler,
	)
	if config.GetString("environment") == "dev" {
		return &Logger{zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
	} else {
		return &Logger{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
	}
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
