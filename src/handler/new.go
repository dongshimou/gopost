package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"model"
	"service"
)

func NewPost(c *gin.Context) {
	var req model.REQNewPost

	err := c.Bind(&req)

	if err != nil {
		log.Print(err)
		return
	}
	err = service.NewPost(&req)
	if err != nil {
		DoResponseFail(c, err)
		return
	}
	DoResponseOK(c, nil)
}
