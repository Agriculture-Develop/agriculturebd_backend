package bizerr

import (
	"context"
	"errors"
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/common/bizcode"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/logs"
	"go.uber.org/zap"
	"runtime"
)

type Biz string

func NewBiz(name string) Biz {
	return Biz(name)
}

func (b Biz) ServiceName() string {
	return string(b)
}

func (b Biz) CodeErr(code bizcode.StatusCode, e ...error) *BizErr {
	if err := errors.Join(e...); err != nil {
		return newBizErr(b.ServiceName(), code, fmt.Errorf("%s: %w", bizcode.ErrMsg[code], err))
	}
	return newBizErr(b.ServiceName(), code, bizcode.ErrMsg[code])
}

type BizErr struct {
	ServiceName string
	Code        bizcode.StatusCode
	Err         error
}

func newBizErr(serviceName string, code bizcode.StatusCode, err error) *BizErr {
	return &BizErr{
		ServiceName: serviceName,
		Code:        code,
		Err:         err,
	}
}

func (e *BizErr) BizMessage() string {
	return fmt.Sprintf("<%s> %s", e.ServiceName, e.Err.Error())
}

func (e *BizErr) Error() string {
	return fmt.Sprintf("[bizerr] service_name: %s, code: %d, error: %s", e.ServiceName, e.Code, e.Err.Error())
}

func (e *BizErr) Log(ctx context.Context, args ...any) *BizErr {
	_, file, line, _ := runtime.Caller(1)
	e.logger(ctx, file, line, fmt.Sprint(args...))

	return e
}

func (e *BizErr) Logf(ctx context.Context, format string, args ...any) *BizErr {
	_, file, line, _ := runtime.Caller(1)
	e.logger(ctx, file, line, fmt.Sprintf(format, args...))

	return e
}

func (e *BizErr) logger(_ context.Context, file string, line int, msg string) {
	logs.Logger.With(
		zap.String("service_name", e.ServiceName),
		zap.String("file", file),
		zap.Int("line", line),
		zap.Int("code", int(e.Code)),
		zap.String("error", e.Err.Error()),
	).Error("%s: %s", msg, e.Err.Error())
}

// TODO : add trace
func (e *BizErr) loggerWithTrace(ctx context.Context, file string, line int, msg string) {
	logs.Logger.With(
		zap.String("service_name", e.ServiceName),
		zap.String("file", file),
		zap.Int("line", line),
		zap.Int("code", int(e.Code)),
		zap.String("error", e.Err.Error()),
	).Error("%s: %s", msg, e.Err.Error())
}
