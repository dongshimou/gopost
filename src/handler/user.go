package handler

import (
	"github.com/gin-gonic/gin"
	"model"
	"service"
)

func SignIn(c *gin.Context) {
	var req model.REQSignin
	var res *model.RESSignIn
	var err error

	req.IP = c.ClientIP()
	//todo ip 限制
	if err = c.Bind(&req); err != nil {
		goto fail
	}

	if res, err = service.SignIn(&req); err != nil {
		goto fail
	}
	setHeaderToken(c, res.Token)
	DoResponseOK(c, res)
	return

fail:
	DoResponseFail(c, err)

}
func SignOut(c *gin.Context) {

	DoResponseOK(c, nil)
	return

}

func SignVerify(c *gin.Context) {

	DoResponseOK(c, nil)
	return
}

func SignUp(c *gin.Context) {

	var req model.REQSignUp
	var res *model.RESSignUp
	var err error

	req.IP = c.ClientIP()
	//todo ip 限制
	if err = c.Bind(&req); err != nil {
		goto fail
	}

	if res, err = service.SignUp(&req); err != nil {
		goto fail
	}
	DoResponseOK(c, res)
	return

fail:
	DoResponseFail(c, err)

}
