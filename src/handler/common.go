package handler

import (
	"base"
	"github.com/gin-gonic/gin"
	"logger"
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

func AuthDecorator(getToken func(string) (*model.User, error), prems ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		c.Request.ParseForm()
		if !logger.DEBUG {
			if token = c.Request.Header.Get("X-User-Token"); len(token) > 0 {
			} else if token = c.Request.Header.Get("USER-TOKEN"); len(token) > 0 {
			} else if token = c.PostForm("user_token"); len(token) > 0 {
			} else if token = c.Query("user_token"); len(token) > 0 {
			} else if token, _ = c.Cookie("user_auth_token"); len(token) > 0 {
			} else {
				logger.Print(utility.ERROR_MSG_AUTH_TOKEN_NOT_EXIST)
				DoResponseFail(c, utility.NewError(utility.ERROR_AUTH_CODE, utility.ERROR_MSG_AUTH_TOKEN_NOT_EXIST))
				c.Abort()
				return
			}
		}
		logger.Print("user token : ", token)
		user, err := getToken(token)
		if err != nil {
			DoResponseFail(c, err)
			c.Abort()
			return
		}
		//验证权限
		if err := utility.VerifyPermission(user.Permission, prems...); err != nil {
			DoResponseFail(c, err)
			c.Abort()
			return
		}
		logger.Print("user info: ", user)
		c.Set("curr_user", user)
		c.Next()
		return
	}
}

func setHeaderToken(c *gin.Context, token string) {
	//cookie:=http.Cookie{
	//Name:"USER-TOKEN",
	//Value:res,
	//Expires:exp,
	//}
	//http.SetCookie(c.Writer,&cookie)
	c.SetCookie("USER-TOKEN", token,
		int(base.GetConfig().Token.TTL),
		"", "", false, true)
}
func getCurrUser(c *gin.Context) (*model.User, error) {
	if a, ok1 := c.Get("curr_user"); ok1 {
		if user, ok2 := a.(*model.User); ok2 {
			return user, nil
		}
	}
	return nil, utility.NewError(utility.ERROR_AUTH_CODE, utility.ERROR_MSG_UNKNOW_USER)
}
