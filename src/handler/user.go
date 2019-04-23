package handler

import (
	"github.com/gin-gonic/gin"
	"protocol"
	"service"
)

func SignIn(c *gin.Context) {
	var (
		req protocol.REQSignin
		res *protocol.RESSignIn
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	req.IP = c.ClientIP()
	//todo ip 限制
	if err = c.Bind(&req); err != nil {
		return
	}

	if res, err = service.SignIn(&req); err != nil {
		return
	}
	setHeaderToken(c, res.Token)
}
func SignOut(c *gin.Context) {
	var (
		err error
	)
	defer func() {
		doResponse(c, err)
	}()
}

func SignVerify(c *gin.Context) {
	var (
		err error
	)
	defer func() {
		doResponse(c, err)
	}()
}

func SignUp(c *gin.Context) {

	var (
		req protocol.REQSignUp
		res interface{}
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	req.IP = c.ClientIP()
	//todo ip 限制
	if err = c.Bind(&req); err != nil {
		return
	}

	if res, err = service.SignUp(&req); err != nil {
		return
	}
}
