// Package handler 包含所有 HTTP 请求处理器
// user.go 处理用户认证相关的 API 请求 (登录、注册、登出等)
package handler

import (
	"github.com/gin-gonic/gin"
	"gopost/src/protocol"
	"gopost/src/service"
)

// ==================== 认证相关处理器 ====================

// SignIn 用户登录
// POST /v1/sign/in
// 请求体: {"username": "xxx", "password": "xxx"}
// 无需认证
func SignIn(c *gin.Context) {
	var (
		req protocol.REQSignin
		res *protocol.RESSignIn
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	// 获取客户端 IP (用于登录限制)
	req.IP = c.ClientIP()
	// TODO: IP 限制防止暴力破解

	// 绑定并验证 JSON 请求体
	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}

	// 调用服务层进行登录验证
	if res, err = service.SignIn(&req); err != nil {
		return
	}
	// 设置响应头中的 Token
	setHeaderToken(c, res.Token)
}

// SignOut 用户登出
// GET /v1/sign/out
// 无需认证
func SignOut(c *gin.Context) {
	var (
		err error
	)
	defer func() {
		doResponse(c, err)
	}()
	// 登出逻辑 (目前为空，可扩展为清除服务端 Session)
}

// SignVerify 验证登录状态
// GET /v1/sign/verify
// 无需认证 (但需要 Token 来验证)
func SignVerify(c *gin.Context) {
	var (
		err error
	)
	defer func() {
		doResponse(c, err)
	}()
	// 验证逻辑由中间件完成，到达这里说明已通过验证
}

// SignUp 用户注册
// POST /v1/sign/up
// 请求体: {"username": "xxx", "password": "xxx", "email": "xxx@xx.com"}
// 无需认证
func SignUp(c *gin.Context) {
	var (
		req protocol.REQSignUp
		res interface{}
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	// 获取客户端 IP (用于注册限制)
	req.IP = c.ClientIP()
	// TODO: IP 限制防止恶意注册

	// 绑定并验证 JSON 请求体
	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}

	// 调用服务层创建用户
	if res, err = service.SignUp(&req); err != nil {
		return
	}
}
