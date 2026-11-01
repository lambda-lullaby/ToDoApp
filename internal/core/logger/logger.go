package core_logger

import (
	"context"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey struct{}

func New(level zapcore.Level, folder string) *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(level)

	if folder != "" {
		if err := os.MkdirAll(folder, 0o755); err != nil {
			panic(err)
		}
		cfg.OutputPaths = append(cfg.OutputPaths, filepath.Join(folder, "app.log"))
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

func ToContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return logger
	}
	return zap.L()
}
