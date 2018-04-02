package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"model"
	"service"
)

func NewReplay(c *gin.Context) {
	var req model.REQNewReplay

	req.Aid = c.Param("aid")
	err := c.Bind(&req)
	if err != nil {
		log.Print(err)
		return
	}
	err = service.NewReplay(&req)
	if err != nil {
		DoResponseFail(c, err)
		return
	}
	DoResponseOK(c, nil)
}
func GetArticle(c *gin.Context) {
	var req model.REQGetArticle

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
func NewPost(c *gin.Context) {
	var req model.REQNewArticle

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
