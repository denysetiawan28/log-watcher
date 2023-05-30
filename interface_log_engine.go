package log_watcher

import (
	"context"
)

type ILogEngine interface {
	Debug(ctx context.Context, message string)
	Info(ctx context.Context, message string)
	Warn(ctx context.Context, message string)
	Error(ctx context.Context, message string)
	Fatal(ctx context.Context, message string)
	Panic(ctx context.Context, message string)
	RequestStarted(ctx context.Context)
	Close() error
}
