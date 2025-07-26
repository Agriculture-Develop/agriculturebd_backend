package resp

import (
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/common/bizcode"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/common/bizerr"
)

func Error(code bizcode.StatusCode, err error) *bizerr.BizErr {
	return bizerr.NewBiz("gateway").CodeErr(code, err)
}

type Response struct {
	StatusCode bizcode.StatusCode `json:"code"`
	StatusMsg  string             `json:"msg"`
	Data       any                `json:"data,omitempty"`
}

func (r *Response) Success(data ...any) {
	r.StatusCode = bizcode.Success
	r.StatusMsg = "success"

	if len(data) > 0 && data[0] != nil {
		r.Data = data[0]
	}
}

func (r *Response) Fail(err *bizerr.BizErr) {
	r.StatusCode = err.Code
	r.StatusMsg = err.BizMessage()
	return
}
