package utility

import (
	"fmt"
	"runtime"
)

const (
	ERROR_OK_CODE = 1000
	ERROR_OK_MSG  = "ok"

	ERROR_UNKNOW_CODE = 1001
	ERROR_UNKNOW_MSG  = "未知错误"

	ERROR_PARSE_CODE = 1002
	ERROR_PARSE_MSG  = "解析错误"

	ERROR_REQUEST_CODE = 1003
	ERROR_REQUEST_MSG  = "请求错误"
)

type InnerError struct {
	Code int
	Msg  string
	Pos  string
}

func (*InnerError) New(code int, args ...interface{}) *InnerError {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	return &InnerError{
		Code: code,
		Msg:  fmt.Sprintf("%v", args...),
		Pos:  fmt.Sprintf("%s %d", file, line),
	}
}
func NewError(code int, args ...interface{}) *InnerError {
	s := InnerError{}
	return s.New(code, args...)
}
func (e *InnerError) Error() string {
	return fmt.Sprintf("Code : %d , Msg : %s ", e.Code, e.Msg)
}
