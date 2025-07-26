package bizcode

import "errors"

type StatusCode int

const (
	Success StatusCode = 200

	Forbidden StatusCode = 300

	BadRequest    StatusCode = 400
	InvalidParams StatusCode = 401
	Unauthorized  StatusCode = 402

	ServerBusy     StatusCode = 500
	RecordNotFound StatusCode = 501
	RecordExist    StatusCode = 502
)

var (
	ErrMsg = map[StatusCode]error{
		Success:   nil,
		Forbidden: errors.New("forbidden"),

		BadRequest:    errors.New("bad request"),
		InvalidParams: errors.New("invalid params"),
		Unauthorized:  errors.New("unauthorized"),

		ServerBusy:     errors.New("server busy"),
		RecordNotFound: errors.New("record not found"),
		RecordExist:    errors.New("record exist"),
	}
)
