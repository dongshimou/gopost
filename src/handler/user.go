package handler

import (
	"github.com/gin-gonic/gin"
	"protocol"
	"service"
)

func SignIn(c *gin.Context) {
	var req protocol.REQSignin
	var res *protocol.RESSignIn
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

	var req protocol.REQSignUp
	var res *protocol.RESSignUp
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
