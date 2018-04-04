package handler

import (
	"github.com/gin-gonic/gin"
	"model"
	"service"
)

func PostReplay(c *gin.Context) {
	var req model.REQNewReplay
	var err error
	if req.CurrUser, err = getCurrUser(c); err != nil {
		goto fail
	}
	if err = c.Bind(&req); err != nil {
		goto fail
	}
	if err = service.NewReplay(&req); err != nil {
		goto fail
	}
	DoResponseOK(c, nil)
	return

fail:
	DoResponseFail(c, err)
}
func GetReplays(c *gin.Context) {
	var req model.REQGetReplays
	var res *model.RESGetReplays
	var err error
	req.Title = c.Param("title")
	if req.CurrUser, err = getCurrUser(c); err != nil {
		goto fail
	}
	if err = c.Bind(&req); err != nil {
		goto fail
	}
	if res, err = service.GetArticleReplays(&req); err != nil {
		goto fail
	}
	DoResponseOK(c, res)
	return

fail:
	DoResponseFail(c, err)
}
func GetUserInfo(c *gin.Context) {
	var req model.REQGetUserInfo
	var res *model.RESGetUserInfo
	var err error
	req.Username = c.Param("username")
	if req.CurrUser, err = getCurrUser(c); err != nil {
		goto fail
	}
	if err = c.Bind(&req); err != nil {
		goto fail
	}
	if res, err = service.GetUserInfo(&req); err != nil {
		goto fail
	}
	DoResponseOK(c, res)
	return

fail:
	DoResponseFail(c, err)

}
func GetArticle(c *gin.Context) {
	var req model.REQGetArticle
	var res *model.RESGetArticle
	var err error
	req.Title = c.Param("title")
	// 等待 gin 的bind url 功能
	if req.CurrUser, err = getCurrUser(c); err != nil {
		goto fail
	}
	if err = c.Bind(&req); err != nil {
		goto fail
	}
	if res, err = service.GetArticle(&req); err != nil {
		goto fail
	}
	DoResponseOK(c, res)
	return

fail:
	DoResponseFail(c, err)
}
func PostArticle(c *gin.Context) {
	var req model.REQNewArticle
	var err error
	if req.CurrUser, err = getCurrUser(c); err != nil {
		goto fail
	}
	if err = c.Bind(&req); err != nil {
		goto fail
	}
	if err = service.PostNewArticle(&req); err != nil {
		goto fail

	}
	DoResponseOK(c, nil)
	return
fail:
	DoResponseFail(c, err)
}
func DelArticle(c *gin.Context) {
	var req model.REQDelArticle
	var err error
	if req.CurrUser, err = getCurrUser(c); err != nil {
		goto fail
	}
	req.Title = c.Param("title")
	if err = c.Bind(&req); err != nil {
		goto fail
	}
	if err = service.DelArticle(&req); err != nil {
		goto fail
	}
	DoResponseOK(c, nil)
	return
fail:
	DoResponseFail(c, err)
}
