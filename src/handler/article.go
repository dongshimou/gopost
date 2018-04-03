package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"model"
	"service"
)

func PostNewReplay(c *gin.Context) {
	var req model.REQNewReplay

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
func GetReplays(c *gin.Context) {
	var req model.REQGetReplays

	req.Title = c.Param("title")
	err := c.Bind(&req)
	if err != nil {
		log.Print(err)
		return
	}
	res, err := service.GetArticleReplays(&req)
	if err != nil {
		DoResponseFail(c, err)
		return
	}
	DoResponseOK(c, res)
}
func GetUserInfo(c *gin.Context) {
	var req model.REQGetUserInfo
	req.Username = c.Param("username")

	err := c.Bind(&req)
	if err != nil {
		log.Print(err)
		return
	}
	res, err := service.GetUserInfo(&req)
	if err != nil {
		DoResponseFail(c, err)
		return
	}
	DoResponseOK(c, res)
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
func PostNewArticle(c *gin.Context) {
	var req model.REQNewArticle
	err := c.Bind(&req)
	if err != nil {
		log.Print(err)
		return
	}
	err = service.PostNewArticle(&req)
	if err != nil {
		DoResponseFail(c, err)
		return
	}
	DoResponseOK(c, nil)
}
