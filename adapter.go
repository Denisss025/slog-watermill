package slogwatermill

import (
	"context"
	"log/slog"

	"github.com/ThreeDotsLabs/watermill"
)

// Adapter config.
type Config struct {
	// Log level (default: LevelDebug).
	Level Level
}

// Adapter implements watermill.LoggerAdapter.
type Adapter struct {
	logger *slog.Logger
	attrs  []slog.Attr
	level  Level
}

var _ = watermill.LoggerAdapter((*Adapter)(nil))

// New creates an instance of Adapter with debug log level and a watermill
// group.
func New(logger *slog.Logger) *Adapter {
	return NewWithConfig(logger, Config{
		Level: LevelDebug,
	})
}

// NewWithConfig creates an instance with a given log level and a watermill
// group.
func NewWithConfig(logger *slog.Logger, config Config) *Adapter {
	return &Adapter{
		level:  config.Level,
		logger: logger.WithGroup("watermill"),
	}
}

func (a Adapter) with(fields watermill.LogFields, attr *slog.Attr) Adapter {
	attrs := make(
		[]slog.Attr, len(a.attrs), len(a.attrs)+len(fields)+1,
	)

	for i := range attrs {
		attrs[i] = a.attrs[i]
	}

	for key, val := range fields {
		attrs = append(attrs, slog.Attr{
			Key:   key,
			Value: slog.AnyValue(val),
		})
	}

	if attr != nil {
		attrs = append(attrs, *attr)
	}

	return Adapter{
		logger: a.logger,
		attrs:  attrs,
		level:  a.level,
	}
}

func (a Adapter) log(level slog.Level, msg string) {
	a.logger.LogAttrs(context.Background(), level, msg, a.attrs...)
}

// Error logs an error.
func (a Adapter) Error(msg string, err error, fields watermill.LogFields) {
	a.with(fields, &slog.Attr{
		Key:   "error",
		Value: slog.AnyValue(err),
	}).log(slog.LevelError, msg)
}

// Info logs some useful information.
func (a Adapter) Info(msg string, fields watermill.LogFields) {
	if a.level > LevelInfo {
		return
	}

	a.with(fields, nil).log(slog.LevelInfo, msg)
}

// Debug logs some debug information.
func (a Adapter) Debug(msg string, fields watermill.LogFields) {
	if a.level > LevelDebug {
		return
	}

	a.with(fields, nil).log(slog.LevelDebug, msg)
}

// Trace logs trace information.
func (a Adapter) Trace(msg string, fields watermill.LogFields) {
	if a.level > LevelTrace {
		return
	}

	a.with(fields, &slog.Attr{
		Key:   "level",
		Value: slog.StringValue("TRACE"),
	}).log(slog.LevelDebug, msg)
}

// With creates a new instance with some extra fields.
func (a Adapter) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return a.with(fields, nil)
}
