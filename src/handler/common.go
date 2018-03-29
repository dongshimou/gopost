package handler

import (
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
	"utility"
)

func doFail(c *gin.Context, code int, msg string) {
	if code < utility.ERROR_OK_CODE {
		c.String(code, msg)
	} else {
		res := model.Response{
			Code: code,
			Msg:  msg,
			Data: nil,
		}
		c.JSON(http.StatusOK, res)
	}
}
func DoResponseFail(c *gin.Context, err error) {

	switch v := err.(type) {
	case *utility.InnerError:
		doFail(c, v.Code, v.Msg)
	default:
		doFail(c, utility.ERROR_UNKNOW_CODE, err.Error())
	}
}
func DoResponseOK(c *gin.Context, data interface{}) {

	var res interface{}

	//组装JSON返回数据
	if data != nil {
		res = model.Response{
			Code: utility.ERROR_OK_CODE,
			Msg:  utility.ERROR_OK_MSG,
			Data: data,
		}
	} else {
		res = model.Response{
			Code: utility.ERROR_OK_CODE,
			Msg:  utility.ERROR_OK_MSG,
			Data: nil,
		}
	}

	c.JSON(http.StatusOK, res)
}
