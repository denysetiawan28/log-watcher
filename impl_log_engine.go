package log_watcher

import (
	"context"
	"go.uber.org/zap"
	"io"
)

type logger struct {
	writers   []io.Writer
	zapLogger *zap.Logger
	level     LogLevel
}

type Field struct {
	Key string
	Val interface{}
}

func NewLogEngine(conf LogConfig) (ILogEngine, error) {
	logEngine := &logger{
		writers: make([]io.Writer, 0),
	}

	logEngine.zapLogger = NewZapLogger(conf, logEngine.writers...)

	return logEngine, nil
}

func (d *logger) Debug(ctx context.Context, message string) {
	//TODO implement me
	zapLogs := []zap.Field{
		zap.String("level", "debug"),
	}

	zapLogs = append(zapLogs, setLogData(ctx, message)...)
	d.zapLogger.Debug("|", zapLogs...)
}

func (d *logger) Info(ctx context.Context, message string) {
	//TODO implement me
	zapLogs := []zap.Field{
		zap.String("level", "info"),
	}

	zapLogs = append(zapLogs, setLogData(ctx, message)...)
	d.zapLogger.Info("|", zapLogs...)
}

func (d *logger) Warn(ctx context.Context, message string) {
	//TODO implement me
	zapLogs := []zap.Field{
		zap.String("level", "warning"),
	}

	zapLogs = append(zapLogs, setLogData(ctx, message)...)
	d.zapLogger.Warn("|", zapLogs...)
}

func (d *logger) Error(ctx context.Context, message string) {
	//TODO implement me
	zapLogs := []zap.Field{
		zap.String("level", "error"),
	}

	zapLogs = append(zapLogs, setLogData(ctx, message)...)
	d.zapLogger.Error("|", zapLogs...)
}

func (d *logger) Fatal(ctx context.Context, message string) {
	//TODO implement me
	zapLogs := []zap.Field{
		zap.String("level", "fatal"),
	}

	zapLogs = append(zapLogs, setLogData(ctx, message)...)
	d.zapLogger.Fatal("|", zapLogs...)
}

func (d *logger) Panic(ctx context.Context, message string) {
	//TODO implement me
	zapLogs := []zap.Field{
		zap.String("level", "panic"),
	}

	zapLogs = append(zapLogs, setLogData(ctx, message)...)
	d.zapLogger.Panic("|", zapLogs...)
}

func (d *logger) RequestStarted(ctx context.Context) {
	//TODO implement me
	zapLogs := []zap.Field{
		zap.String("level", "info"),
	}

	zapLogs = append(zapLogs, setStartRequestLogData(ctx)...)
	d.zapLogger.Info("|", zapLogs...)
}

func (d *logger) Close() error {
	//TODO implement me
	panic("implement me")
}
