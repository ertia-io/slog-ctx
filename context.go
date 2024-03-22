package slogx

import (
	"context"
	"io"
	"log/slog"
)

type contextKey struct{}

var nilLogger = slog.New(slog.NewTextHandler(io.Discard, nil))

// Ctx retrieves an attached slog.Logger from the context.
// If no logger is attached, a logger with io.Discard will be returned.
func Ctx(ctx context.Context) *slog.Logger {
	raw := ctx.Value(contextKey{})
	if raw == nil {
		return nilLogger
	}

	logger, ok := raw.(*slog.Logger)
	if !ok {
		return nilLogger
	}

	return logger
}

// WithContext attaches a slog.Logger to the provided context
func WithContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

// With attaches given attributes to the slog.Logger found inside the context.
// If there is no logger on the context, it will attach a logger with io.Discard.
func With(ctx context.Context, args ...any) context.Context {
	return WithContext(ctx, Ctx(ctx).With(args...))
}
