// Package handler 包含所有 HTTP 请求处理器
// article.go 处理文章和评论相关的 API 请求
package handler

import (
	"github.com/gin-gonic/gin"
	"gopost/src/logger"
	"gopost/src/protocol"
	"gopost/src/service"
)

// ==================== 评论相关处理器 ====================

// PostReplay 创建新评论
// POST /v1/replay
// 需要: 评论创建权限
func PostReplay(c *gin.Context) {
	var (
		req protocol.REQNewReplay
		err error
	)
	defer func() {
		doResponse(c, err)
	}()
	if req.CurrUser, err = getCurrUser(c); err != nil {
		return
	}
	// 绑定并验证 JSON 请求体
	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}
	// 获取客户端 IP
	req.IpAddress = c.ClientIP()
	// 调用服务层创建评论
	if err = service.PostReplay(&req); err != nil {
		return
	}
}

// GetReplays 获取评论列表
// GET /v1/replay/:title
func GetReplays(c *gin.Context) {
	var (
		req protocol.REQGetReplays
		res interface{}
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	req.Title = c.Param("title")
	if req.CurrUser, err = getCurrUser(c); err != nil {
		return
	}
	if err = c.Bind(&req); err != nil {
		return
	}
	if res, err = service.GetArticleReplays(&req); err != nil {
		return
	}
}
func DelReplays(c *gin.Context) {
	var (
		req protocol.REQDelReplays
		err error
	)
	defer func() {
		doResponse(c, err)
	}()
	if req.CurrUser, err = getCurrUser(c); err != nil {
		return
	}
	req.Rid = c.Param("rid")
	if err = service.DelArticleReplay(&req); err != nil {
		return
	}
}
func GetUserInfo(c *gin.Context) {
	var (
		req protocol.REQGetUserInfo
		res interface{}
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	req.Username = c.Param("username")
	if req.CurrUser, err = getCurrUser(c); err != nil {
		return
	}
	if err = c.Bind(&req); err != nil {
		return
	}
	if res, err = service.GetUserInfo(&req); err != nil {
		return
	}
}
func GetArticles(c *gin.Context) {
	var (
		req protocol.REQGetArticles
		res interface{}
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	if err = c.Bind(&req); err != nil {
		return
	}
	if res, err = service.GetArticles(&req); err != nil {
		return
	}
}

// SearchArticles 搜索文章
// GET /v1/article/search?keyword=xxx
func SearchArticles(c *gin.Context) {
	var (
		req protocol.REQSearchArticle
		res interface{}
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	if err = c.Bind(&req); err != nil {
		return
	}
	if res, err = service.SearchArticles(&req); err != nil {
		return
	}
}

// GetArticlesByTag 按标签获取文章
// GET /v1/article/tag?tag=xxx&size=10
func GetArticlesByTag(c *gin.Context) {
	var (
		req protocol.REQGetArticlesByTag
		res interface{}
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	if err = c.Bind(&req); err != nil {
		return
	}
	if res, err = service.GetArticlesByTag(&req); err != nil {
		return
	}
}

func StatIp(c *gin.Context) {
	ip := c.ClientIP()
	if err := service.StatIp(ip); err != nil {
		logger.Error(err)
	}
}
func GetStat(c *gin.Context) {
	var (
		req protocol.REQGetStat
		res interface{}
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	if err = c.Bind(&req); err != nil {
		return
	}
	if res, err = service.GetStat(&req); err != nil {
		return
	}
}
func GetArticle(c *gin.Context) {
	var (
		req protocol.REQGetArticle
		res interface{}
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	req.Title = c.Param("title")
	// 等待 gin 的bind url 功能
	if req.CurrUser, err = getCurrUser(c); err != nil {
		return
	}
	if err = c.Bind(&req); err != nil {
		return
	}
	if res, err = service.GetArticle(&req); err != nil {
		return
	}
}
func CreateArticle(c *gin.Context) {
	var (
		req protocol.REQNewArticle
		err error
	)
	defer func() {
		doResponse(c, err)
	}()
	if req.CurrUser, err = getCurrUser(c); err != nil {
		return
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}
	if err = service.CreateArticle(&req); err != nil {
		return
	}
}
func UpdateArticle(c *gin.Context) {
	var (
		req protocol.REQUpdateArticle
		err error
	)
	defer func() {
		doResponse(c, err)
	}()
	if req.CurrUser, err = getCurrUser(c); err != nil {
		return
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}
	req.OldTitle = c.Param("oldtitle")
	if err = service.UpdateArticle(&req); err != nil {
		return
	}
}
func DelArticle(c *gin.Context) {
	var (
		req protocol.REQDelArticle
		err error
	)
	defer func() {
		doResponse(c, err)
	}()
	if req.CurrUser, err = getCurrUser(c); err != nil {
		return
	}
	req.Title = c.Param("title")
	if err = c.Bind(&req); err != nil {
		return
	}
	if err = service.DelArticle(&req); err != nil {
		return
	}
}
func GetTags(c *gin.Context) {
	var (
		req protocol.REQGetTags
		res interface{}
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	req.Title = c.Param("title")
	if err = c.Bind(&req); err != nil {
		return
	}
	if res, err = service.GetTags(&req); err != nil {
		return
	}
}
func GetAllTags(c *gin.Context) {
	var (
		res interface{}
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	if res, err = service.GetAllTags(); err != nil {
		return
	}
}
