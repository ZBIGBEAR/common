package logger

import (
	"context"
	"fmt"
	"github.com/ZBIGBEAR/common/consts"
	"github.com/k0kubun/pp/v3"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"os"
	"time"
)

var globalZerolog = zerolog.New(os.Stderr).With().Logger()

type Req struct {
	RequestMeta

	CustomTag string `json:"serviceName"`       // required,服务名
	Level     string `json:"level"`             // required,日志类型[INFO，WRAN，ERROR]
	Date      string `json:"date"`              // required,日志产生时间
	Message   string `json:"message,omitempty"` // optional,日志信息
}

type RequestMeta struct {
	Scheme      string `json:"scheme"`             // required,接口协议 [http,grpc]
	RequestId   string `json:"requestId"`          // required,由istio-gateway生成的请求id,从header中x-request-id获取，需要一直往后传递，追踪完整链路
	Transaction string `json:"transaction"`        // required,请求的接口/方法名
	Path        string `json:"path"`               // optional,请求path
	Duration    int64  `json:"duration"`           // required,本次请求的时长
	Method      string `json:"method,omitempty"`   // optional,http 方法[GET,POST...]
	Query       string `json:"query,omitempty"`    // optional,http query 字符串
	Body        []byte `json:"body,omitempty"`     // optional,http post/put body json字符串
	ClientIP    string `json:"clientIp,omitempty"` // optional,来源ip
	Token       string `json:"token,omitempty"`    // optional,本次请求的用户token
	Status      int    `json:"status,omitempty"`   // optional,状态码
	UserId      string `json:"userId,omitempty"`   // optional,本次请求的的userId
	Header      []byte `json:"header,omitempty"`   // optional,本次请求的header
}

type log struct {
	notify      Notify
	serviceName string
}

type Option func(l *log)

func New(options ...Option) Logger {
	l := &log{}
	for i := range options {
		options[i](l)
	}
	return l
}

func (l *log) baseLog(ctx context.Context, zLog *zerolog.Event) *zerolog.Event {
	if ctx == nil {
		ctx = context.TODO()
	}
	logCtx := zLog.Time(consts.LblDate, time.Now()).Str(consts.LblServiceName, l.serviceName)
	req := ctx.Value(consts.ReqObjKey)
	if req == nil {
		return logCtx
	}

	reqObj, ok := req.(Req)
	if !ok {
		return logCtx
	}
	if reqObj.UserId != "" {
		logCtx = logCtx.Str(consts.LblUserId, reqObj.UserId)
	}

	return logCtx.Str(consts.LblScheme, reqObj.Scheme).
		Str(consts.LblRequestId, reqObj.RequestId).
		Str(consts.LblTransaction, reqObj.Transaction).
		Str(consts.LblMethod, reqObj.Method).
		Str(consts.LblQuery, reqObj.Query).
		Str(consts.LblClientIP, reqObj.ClientIP).
		Interface(consts.LblToken, ctx.Value(consts.LblToken))
}

func (l *log) Debugf(ctx context.Context, format string, attr ...interface{}) {
	l.Debug(ctx, fmt.Sprintf(format, attr...))
}

func (l *log) Debug(ctx context.Context, message string) {
	l.Debugw(ctx, message)
}

func (l *log) Debugw(ctx context.Context, message string, keysAndValues ...interface{}) {
	l.baseLog(ctx, globalZerolog.Debug()).
		Interface(consts.LblMsg, serialKeysAndValues(message, keysAndValues...)).Send()
}

func (l *log) Infof(ctx context.Context, format string, attr ...interface{}) {
	l.Info(ctx, fmt.Sprintf(format, attr...))
}

func (l *log) Info(ctx context.Context, message string) {
	l.Infow(ctx, message)
}

func (l *log) Infow(ctx context.Context, message string, keysAndValues ...interface{}) {
	l.baseLog(ctx, globalZerolog.Info()).Interface(consts.LblMsg, serialKeysAndValues(message, keysAndValues...)).Send()
}

func (l *log) Warnf(ctx context.Context, format string, attr ...interface{}) {
	l.Warn(ctx, fmt.Sprintf(format, attr...))
}

func (l *log) Warn(ctx context.Context, message string) {
	l.Warnw(ctx, message)
}

func (l *log) Warnw(ctx context.Context, message string, keysAndValues ...interface{}) {
	l.baseLog(ctx, globalZerolog.Warn()).
		Interface(consts.LblMsg, serialKeysAndValues(message, keysAndValues...)).Send()
}

func (l *log) Errorf(ctx context.Context, format string, attr ...interface{}) {
	l.Error(ctx, fmt.Sprintf(format, attr...))
}

func (l *log) Error(ctx context.Context, message string) {
	l.Errorw(ctx, message)
}

func (l *log) Errorw(ctx context.Context, message string, keysAndValues ...interface{}) {
	// 发送通知
	kv := serialKeysAndValues(message, keysAndValues...)
	l.sendMessage(ctx, message, kv)

	l.baseLog(ctx, globalZerolog.Error()).
		Interface(consts.LblMsg, kv).
		Send()
}

func (l *log) Fatalf(ctx context.Context, format string, attr ...interface{}) {
	l.Fatal(ctx, fmt.Sprintf(format, attr...))
}

func (l *log) Fatal(ctx context.Context, message string) {
	l.Fatalw(ctx, message)
}

func (l *log) Fatalw(ctx context.Context, message string, keysAndValues ...interface{}) {
	// 发送通知
	kv := serialKeysAndValues(message, keysAndValues...)
	l.sendMessage(ctx, message, kv)

	l.baseLog(ctx, globalZerolog.Fatal()).
		Interface(consts.LblMsg, kv).Send()
}

func serialKeysAndValues(msg string, args ...interface{}) map[string]interface{} {
	fields := map[string]interface{}{
		consts.LblMessage: msg,
	}
	for i := 0; i < len(args); {
		if i == len(args)-1 {
			fields["ignored"] = args[i]
			break
		}
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			fields[fmt.Sprintf("%s", key)] = val
		} else {
			fields[keyStr] = val
		}
		i += 2
	}
	return fields
}

type stackTracer interface {
	Error() string
	StackTrace() errors.StackTrace
}

func (l *log) sendMessage(ctx context.Context, message string, data map[string]interface{}) {
	if l.notify == nil {
		return
	}

	// 异步发送
	go func() {
		msg := map[string]interface{}{}
		pp.Default.SetColoringEnabled(false)
		pp.PrintMapTypes = false
		var traceInfo string
		for k, v := range data {
			msg[k] = v
			switch vn := v.(type) {
			case stackTracer:
				msg[k] = vn.Error()
				for _, f := range vn.StackTrace() {
					traceInfo += fmt.Sprintf("%+s:%d\n", f, f)
				}
			case error:
				msg[k] = vn.Error()
			}
		}
		delete(msg, "message")
		d := pp.Sprint(msg)
		if traceInfo != "" {
			d = fmt.Sprintf("%s\n\n=== stacktrace ===\n%s", d, traceInfo)
		}
		if err := l.notify.Notify(ctx, fmt.Sprintf("[%s] ERROR MESSAGE:%s\n%s", l.serviceName, message, d)); err != nil {
			fmt.Printf("log.sendMessage. err:%v\n", err)
		}
	}()
}

// default global logger
var (
	_default Logger
)

func init() {
	options := []Option{
		WithServiceName("go-template"),
	}

	feiShuWebhook := os.Getenv("FEISHU_WEBHOOK")
	if feiShuWebhook != "" {
		options = append(options, WithFeiShuNotify(feiShuWebhook))
	}
	_default = New(options...)
}

func Debugf(ctx context.Context, format string, attr ...interface{}) {
	_default.Debugf(ctx, format, attr...)
}

func Debug(ctx context.Context, message string) {
	_default.Debug(ctx, message)
}

func Debugw(ctx context.Context, message string, keysAndValues ...interface{}) {
	_default.Debugw(ctx, message, keysAndValues...)
}

func Infof(ctx context.Context, format string, attr ...interface{}) {
	_default.Infof(ctx, format, attr...)
}

func Info(ctx context.Context, message string) {
	_default.Info(ctx, message)
}

func Infow(ctx context.Context, message string, keysAndValues ...interface{}) {
	_default.Infow(ctx, message, keysAndValues...)
}

func Warnf(ctx context.Context, format string, attr ...interface{}) {
	_default.Warnf(ctx, format, attr...)
}

func Warn(ctx context.Context, message string) {
	_default.Warn(ctx, message)
}

func Warnw(ctx context.Context, message string, keysAndValues ...interface{}) {
	_default.Warnw(ctx, message, keysAndValues)
}

func Errorf(ctx context.Context, format string, attr ...interface{}) {
	_default.Errorf(ctx, format, attr...)
}

func Error(ctx context.Context, message string) {
	_default.Error(ctx, message)
}

func Errorw(ctx context.Context, message string, keysAndValues ...interface{}) {
	_default.Errorw(ctx, message, keysAndValues...)
}

func Fatalf(ctx context.Context, format string, attr ...interface{}) {
	_default.Fatalf(ctx, format, attr...)
}

func Fatal(ctx context.Context, message string) {
	_default.Fatal(ctx, message)
}

func Fatalw(ctx context.Context, message string, keysAndValues ...interface{}) {
	_default.Fatalw(ctx, message, keysAndValues...)
}
