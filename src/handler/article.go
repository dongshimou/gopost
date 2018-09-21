package handler

import (
	"github.com/gin-gonic/gin"
	"model"
	"service"
	"logger"
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
	req.IpAddress = c.ClientIP()
	if err = service.PostReplay(&req); err != nil {
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
func DelReplays(c *gin.Context) {
	var req model.REQDelReplays
	var err error

	if req.CurrUser, err = getCurrUser(c); err != nil {
		goto fail
	}
	req.Rid = c.Param("rid")
	if err = service.DelArticleReplay(&req); err != nil {
		goto fail
	}
	DoResponseOK(c, nil)
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
func GetArticles(c *gin.Context) {
	var req model.REQGetArticles
	var res *model.RESGetArticles
	var err error
	if err = c.Bind(&req); err != nil {
		goto fail
	}
	if res, err = service.GetArticles(&req); err != nil {
		goto fail
	}
	DoResponseOK(c, res)
	return
fail:
	DoResponseFail(c, err)
}
func StatIp(c *gin.Context){
	ip:=c.ClientIP()
	if err:=service.StatIp(ip);err!=nil{
		logger.Error(err)
	}
}
func GetStat(c *gin.Context){
	var req model.REQGetStat
	var res *model.RESGetStat
	var err error
	if err=c.Bind(&req);err!=nil{
		goto fail
	}
	if res,err=service.GetStat(&req);err!=nil{
		goto fail
	}
	DoResponseOK(c,res)
	return
fail:
	DoResponseFail(c,err)
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
func CreateArticle(c *gin.Context) {
	var req model.REQNewArticle
	var err error
	if req.CurrUser, err = getCurrUser(c); err != nil {
		goto fail
	}
	if err = c.Bind(&req); err != nil {
		goto fail
	}
	if err = service.CreateArticle(&req); err != nil {
		goto fail

	}
	DoResponseOK(c, nil)
	return
fail:
	DoResponseFail(c, err)
}
func UpdateArticle(c *gin.Context) {
	var req model.REQUpdateArticle
	var err error
	if req.CurrUser, err = getCurrUser(c); err != nil {
		goto fail
	}
	if err = c.Bind(&req); err != nil {
		goto fail
	}
	req.OldTitle = c.Param("oldtitle")
	if err = service.UpdateArticle(&req); err != nil {
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
func GetTags(c *gin.Context) {
	var res *model.RESGetTags
	var err error
	if res, err = func() (*model.RESGetTags, error) {
		var req model.REQGetTags
		req.Title = c.Param("title")
		if err := c.Bind(&req); err != nil {
			return nil, err
		}
		return service.GetTags(&req)
	}(); err != nil {
		DoResponseFail(c, err)
	}
	DoResponseOK(c, res)
}
func GetAllTags(c *gin.Context) {
	res, err := service.GetAllTags()
	if err != nil {
		DoResponseFail(c, err)
	}
	DoResponseOK(c, res)
}
