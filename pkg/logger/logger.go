package logger

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zbitech/common/pkg/rctx"
)

type LogLevel int8

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
	TraceLevel LogLevel = -1
)

func SetGlobalLevel(level LogLevel) {
	zerolog.SetGlobalLevel(zerolog.Level(level))
}

func LogComponentTime(ctx context.Context) {
	_startTime := ctx.Value(rctx.StartTime)

	if _startTime != nil {
		startTime := _startTime.(time.Time)
		runningTime := time.Now().Sub(startTime)

		Infof(ctx, "RunningTime %d ms", runningTime.Milliseconds())
	}
}

func Trace(ctx context.Context, message string) {
	Msg(TraceLevel, ctx, message)
}

func Tracef(ctx context.Context, message string, params ...interface{}) {
	Msgf(TraceLevel, ctx, message, params...)
}

func Debug(ctx context.Context, message string) {
	Msg(DebugLevel, ctx, message)
}

func Debugf(ctx context.Context, message string, params ...interface{}) {
	Msgf(DebugLevel, ctx, message, params...)
}

func Info(ctx context.Context, message string) {
	Msg(InfoLevel, ctx, message)
}

func Infof(ctx context.Context, message string, params ...interface{}) {
	Msgf(InfoLevel, ctx, message, params...)
}

func Warn(ctx context.Context, message string) {
	Msg(WarnLevel, ctx, message)
}

func Warnf(ctx context.Context, message string, params ...interface{}) {
	Msgf(WarnLevel, ctx, message, params...)
}

func Error(ctx context.Context, message string) {
	Msg(ErrorLevel, ctx, message)
}

func Errorf(ctx context.Context, message string, params ...interface{}) {
	Msgf(ErrorLevel, ctx, message, params...)
}

func Fatal(ctx context.Context, message interface{}) {
	event := log.Fatal()
	event = addEventParams(event, ctx)
	event.Msgf("%s", message)
}

func Fatalf(ctx context.Context, message string, params ...interface{}) {
	Msgf(FatalLevel, ctx, message, params...)
}

func Panic(ctx context.Context, message string) {
	Msg(PanicLevel, ctx, message)
}

func Panicf(ctx context.Context, message string, params ...interface{}) {
	Msgf(PanicLevel, ctx, message, params...)
}

func getEvent(level LogLevel) *zerolog.Event {
	switch level {
	case TraceLevel:
		return log.Trace()
	case DebugLevel:
		return log.Debug()
	case InfoLevel:
		return log.Info()
	case WarnLevel:
		return log.Warn()
	case ErrorLevel:
		return log.Error()
	case FatalLevel:
		return log.Fatal()
	case PanicLevel:
		return log.Panic()
	default:
		return log.Info()
	}
}

func addEventParam(event *zerolog.Event, ctx context.Context, param rctx.ContextParam, default_value string) *zerolog.Event {

	key := string(param)
	value := ctx.Value(param)
	if value == nil {
		value = default_value
	}

	return event.Str(key, value.(string))
}

func addEventParams(event *zerolog.Event, ctx context.Context) *zerolog.Event {

	n_event := addEventParam(event, ctx, rctx.RequestId, "xxx")
	n_event = addEventParam(n_event, ctx, rctx.Component, "ZBI")

	return n_event
}

func Msg(level LogLevel, ctx context.Context, message string) {
	var event *zerolog.Event = getEvent(level)
	event = addEventParams(event, ctx)
	event.Msg(message)
}

func Msgf(level LogLevel, ctx context.Context, message string, params ...interface{}) {
	var event *zerolog.Event = getEvent(level)
	event = addEventParams(event, ctx)
	event.Msgf(message, params...)
}
