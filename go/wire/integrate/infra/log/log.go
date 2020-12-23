package log

import (
	"context"

	"github.com/google/wire"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
}

// New zap logger
func New() (*zap.Logger, func()) {
	// logger, err := zap.NewProduction()
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(errors.Wrap(err, "New zap logger failed."))
	}

	return logger, func() { logger.Sync() }
}

// Extract zap.Logger from context
func CtxZap() func(context.Context) *zap.Logger {
	return func(ctx context.Context) *zap.Logger {
		return ctxzap.Extract(ctx)
	}
}

// Extract zap.Logger from context and use its SugaredLogger
func CtxLogger() func(context.Context) Logger {
	return func(ctx context.Context) Logger {
		return CtxZap()(ctx).Sugar()
	}
}

var ProviderSet = wire.NewSet(New, CtxZap, CtxLogger)
