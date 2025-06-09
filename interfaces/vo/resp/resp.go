package resp

import (
	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
)

type Response struct {
	StatusCode respCode.StatusCode `json:"code"`
	StatusMsg  string              `json:"msg,omitempty"` // omitempty如果字段为空则忽略显示
	Data       interface{}         `json:"data,omitempty"`
}

// Msg returns the message of the resp
func (r *Response) Msg() string {
	if m, ok := respCode.Msg[r.StatusCode]; ok {
		return m
	}
	return ""
}

// SetNoData prepares the response without data
func (r *Response) SetNoData(code respCode.StatusCode) {
	r.StatusCode = code
	r.StatusMsg = r.Msg()
}

// SetWithData prepares the response with data
func (r *Response) SetWithData(code respCode.StatusCode, data interface{}) {
	r.StatusCode = code
	r.StatusMsg = r.Msg()
	r.Data = data
}

// GetMsg returns the message of the resp without Response type
func GetMsg(code respCode.StatusCode) string {
	if msg, ok := respCode.Msg[code]; ok {
		return msg
	}
	return ""
}
