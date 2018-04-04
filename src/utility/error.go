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

	ERROR_AUTH_CODE = 1004
	ERROR_AUTH_MSG  = "权限错误"

	ERROR_MSG_UNKNOW_USER = "未知用户"
)

type InnerError struct {
	Code int
	Msg  string
	Pos  string
}

func (*InnerError) New(code int, msg string, args ...interface{}) *InnerError {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	msg = fmt.Sprintf("%s ", msg)
	for _, v := range args {
		msg += fmt.Sprintf("%v", v)
	}
	return &InnerError{
		Code: code,
		Msg:  msg,
		Pos:  fmt.Sprintf("%s %d", file, line),
	}
}
func NewError(code int, msg string, args ...interface{}) *InnerError {
	s := InnerError{}
	return s.New(code, msg, args...)
}
func Wrap(err error, code int, msg string, args ...interface{}) *InnerError {
	if err == nil {
		return NewError(code, msg, args...)
	}
	v, ok := err.(*InnerError)
	if ok {
		args = append(args, msg)
		return NewError(v.Code, v.Msg, args...)
	} else {
		return NewError(code, msg, args...)
	}
}
func (e *InnerError) Error() string {
	return fmt.Sprintf("Code : %d , Msg : %s ", e.Code, e.Msg)
}
