package logger

import (
	"context"
	"log/slog"
	"os"
)

const (
	DebugLogLvl = "DEBUG"
	WarnLogLvl  = "WARN"
	ErrLogLvl   = "ERROR"
)

type (
	Logger interface {
		Debug(msg string, args ...any)
		DebugCtx(ctx context.Context, msg string, args ...any)

		Info(msg string, args ...any)
		InfoCtx(ctx context.Context, msg string, args ...any)

		Warn(msg string, args ...any)
		WarnCtx(ctx context.Context, msg string, args ...any)

		Error(msg string, args ...any)
		ErrorCtx(ctx context.Context, msg string, args ...any)
	}

	LoggerImpl struct {
		l *slog.Logger
	}

	LoggerCfg struct {
		Lvl string
	}
)

func NewLoggerImpl(cfg *LoggerCfg) Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: getLogLvl(cfg.getLogLvl())}))

	return &LoggerImpl{l: logger}
}

func (l *LoggerImpl) Debug(msg string, args ...any) {
	l.l.Debug(msg, args...)
}

func (l *LoggerImpl) DebugCtx(ctx context.Context, msg string, args ...any) {
	l.l.DebugContext(ctx, msg, args...)
}

func (l *LoggerImpl) Info(msg string, args ...any) {
	l.l.Info(msg, args...)
}

func (l *LoggerImpl) InfoCtx(ctx context.Context, msg string, args ...any) {
	l.l.InfoContext(ctx, msg, args...)
}

func (l *LoggerImpl) Warn(msg string, args ...any) {
	l.l.Warn(msg, args...)
}

func (l *LoggerImpl) WarnCtx(ctx context.Context, msg string, args ...any) {
	l.l.WarnContext(ctx, msg, args...)
}

func (l *LoggerImpl) Error(msg string, args ...any) {
	l.l.Error(msg, args...)
}

func (l *LoggerImpl) ErrorCtx(ctx context.Context, msg string, args ...any) {
	l.l.ErrorContext(ctx, msg, args...)
}

func (l *LoggerCfg) getLogLvl() string {
	if l != nil {
		return l.Lvl
	}
	return ""
}

func getLogLvl(logLvl string) slog.Level {
	switch logLvl {
	case ErrLogLvl:
		return slog.LevelError
	case WarnLogLvl:
		return slog.LevelWarn
	case DebugLogLvl:
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}
