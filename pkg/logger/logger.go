package logger

import (
	"context"
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	LogFile    string `mapstructure:"log_file"`
	ErrorFile  string `mapstructure:"error_file"`
	MaxSize    int    `mapstructure:"max_size"` // megabytes
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"` // days
	Compress   bool   `mapstructure:"compress"`
	Level      string `mapstructure:"level"`
}

type CustomHandler struct {
	appHandler slog.Handler
	errHandler slog.Handler
}

func (h *CustomHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.appHandler.Enabled(ctx, level)
}

func (h *CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	if r.Level >= slog.LevelError {
		_ = h.errHandler.Handle(ctx, r)
	}
	return h.appHandler.Handle(ctx, r)
}

func (h *CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &CustomHandler{
		appHandler: h.appHandler.WithAttrs(attrs),
		errHandler: h.errHandler.WithAttrs(attrs),
	}
}

func (h *CustomHandler) WithGroup(name string) slog.Handler {
	return &CustomHandler{
		appHandler: h.appHandler.WithGroup(name),
		errHandler: h.errHandler.WithGroup(name),
	}
}

func InitLogger(cfg *Config) {
	var logLevel slog.Level
	switch cfg.Level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	appLog := &lumberjack.Logger{
		Filename:   cfg.LogFile,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	errLog := &lumberjack.Logger{
		Filename:   cfg.ErrorFile,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	appWriter := io.MultiWriter(os.Stdout, appLog)
	errWriter := io.MultiWriter(os.Stderr, errLog)

	handler := &CustomHandler{
		appHandler: slog.NewJSONHandler(appWriter, opts),
		errHandler: slog.NewJSONHandler(errWriter, opts),
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
