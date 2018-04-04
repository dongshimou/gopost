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

	ERROR_CONVER_CODE = 1005
	ERROR_CONVER_MSG  = "类型转换错误"

	ERROR_MSG_UNKNOW_USER       = "未知用户"
	ERROR_MSG_EXIST_USERNAME    = "已存在的用户名"
	ERROR_MSG_ERROR_USERNAME    = "错误的用户名"
	ERROR_MSG_LENGTH_USERNAME   = "用户名长度应为6-30个字符"
	ERROR_MSG_ERROR_PASSWORD    = "密码错误"
	ERROR_MSG_ERROR_TOKEN_SIGN  = "认证签名错误"
	ERROR_MSG_ERROR_SIGN_METHOD = "错误的签名方式"
	ERROR_MSG_AUTH_TIMEOUT      = "认证超时,请重新登录"

	ERROR_MSG_AUTH_TOKEN_KNOW_ERROR     = "解析token失败"
	ERROR_MSG_AUTH_TOKEN_EXP_ERROR      = "从token获取时间失败"
	ERROR_MSG_AUTH_TOKEN_USERNAME_ERROR = "从token获取用户失败"
	ERROR_MSG_AUTH_TOKEN_NOT_EXIST      = "未找到token"
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
	msg = fmt.Sprintf("%s", msg)
	for _, v := range args {
		msg += fmt.Sprintf(" %v", v)
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
		return NewError(v.Code, v.Msg, args...)
	} else {
		return NewError(code, msg, args...)
	}
}
func (e *InnerError) Error() string {
	return fmt.Sprintf("Code : %d , Msg : %s ", e.Code, e.Msg)
}
