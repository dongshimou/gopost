package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"model"
	"service"
)

func Signin(c *gin.Context) {
	var req model.REQLogin

	err := c.Bind(&req)

	if err != nil {
		log.Print(err)
		return
	}
	res, err := service.Login(&req)
	if err != nil {
		DoResponseFail(c, err)
		return
	}
	DoResponseOK(c, res)
}
func SignOut(c *gin.Context) {

	DoResponseOK(c, nil)
}

func SignUp(c *gin.Context) {

}
