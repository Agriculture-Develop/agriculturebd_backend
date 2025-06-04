package resp

type Response struct {
	StatusCode int64       `json:"code"`
	StatusMsg  string      `json:"msg,omitempty"` // omitempty如果字段为空则忽略显示
	Data       interface{} `json:"data"`
}

// Msg returns the message of the response
func (r *Response) Msg() string {
	if m, ok := apiCode.Msg[r.StatusCode]; ok {
		return m
	}
	return ""
}

// SetNoData prepares the response without data
func (r *Response) SetNoData(code int64) {
	r.StatusCode = code
	r.StatusMsg = r.Msg()
}

// SetWithData prepares the response with data
func (r *Response) SetWithData(code int64, data interface{}) {
	r.StatusCode = code
	r.StatusMsg = r.Msg()
	r.Data = data
}
