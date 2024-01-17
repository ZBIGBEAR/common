package logger

import "context"

type Logger interface {
	Debugf(ctx context.Context, format string, attr ...interface{})
	Debug(ctx context.Context, message string)
	Debugw(ctx context.Context, message string, keysAndValues ...interface{})
	Infof(ctx context.Context, format string, attr ...interface{})
	Info(ctx context.Context, message string)
	Infow(ctx context.Context, message string, keysAndValues ...interface{})
	Warnf(ctx context.Context, format string, attr ...interface{})
	Warn(ctx context.Context, message string)
	Warnw(ctx context.Context, message string, keysAndValues ...interface{})
	Errorf(ctx context.Context, format string, attr ...interface{})
	Error(ctx context.Context, message string)
	Errorw(ctx context.Context, message string, keysAndValues ...interface{})
	Fatalf(ctx context.Context, format string, attr ...interface{})
	Fatal(ctx context.Context, message string)
	Fatalw(ctx context.Context, message string, keysAndValues ...interface{})
}

type Notify interface {
	Notify(ctx context.Context, msg string) error
}
