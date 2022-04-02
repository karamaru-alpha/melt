package app

import (
	"context"

	"go.uber.org/zap"

	"github.com/karamaru-alpha/melt/pkg/merrors"
)

const logName = "app"

type Logger interface {
	Debug(ctx context.Context, msg string, fields ...zap.Field)
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
}

type logger struct {
	*zap.Logger
	isLocal bool
}

var appLogger Logger

func New(isLocal bool) (Logger, error) {
	var config zap.Config
	if isLocal {
		config = zap.NewDevelopmentConfig()
	} else {
		// TODO: 本番環境のログ基盤
		config = zap.NewDevelopmentConfig()
	}

	l, err := config.Build()
	if err != nil {
		return nil, merrors.Stack(err)
	}
	return &logger{Logger: l.Named(logName), isLocal: isLocal}, nil
}

func GetLogger() Logger {
	if appLogger == nil {
		l, _ := New(false)
		l.Error(context.Background(), "not initial set AppLogger")
		return l
	}
	return appLogger
}

func (l *logger) Debug(_ context.Context, msg string, fields ...zap.Field) {
	l.Logger.Debug(msg, fields...)
}

func (l *logger) Info(_ context.Context, msg string, fields ...zap.Field) {
	l.Logger.Info(msg, fields...)
}

func (l *logger) Error(_ context.Context, msg string, fields ...zap.Field) {
	l.Logger.Error(msg, fields...)
}
