package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"model"
	"service"
)

func GetPost(c *gin.Context) {
	var req model.REQGetPost

	req.Title = c.Param("title")
	// 等待 gin 的bind url 功能
	err := c.Bind(&req)
	if err != nil {
		log.Print(err)
		return
	}
	res, err := service.GetPost(&req)
	if err != nil {
		DoResponseFail(c, err)
		return
	}
	DoResponseOK(c, res)
}
