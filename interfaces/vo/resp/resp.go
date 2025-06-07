package resp

type Response struct {
	StatusCode StatusCode  `json:"code"`
	StatusMsg  string      `json:"msg,omitempty"` // omitempty如果字段为空则忽略显示
	Data       interface{} `json:"data"`
}

// Msg returns the message of the resp
func (r *Response) Msg() string {
	if m, ok := Msg[r.StatusCode]; ok {
		return m
	}
	return ""
}

// GetMsg returns the message of the resp without Response type
func GetMsg(code StatusCode) string {
	if msg, ok := Msg[code]; ok {
		return msg
	}
	return ""
}

// SetNoData prepares the response without data
func (r *Response) SetNoData(code StatusCode) {
	r.StatusCode = code
	r.StatusMsg = r.Msg()
}

// SetWithData prepares the response with data
func (r *Response) SetWithData(code StatusCode, data interface{}) {
	r.StatusCode = code
	r.StatusMsg = r.Msg()
	r.Data = data
}
