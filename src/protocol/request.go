// Package protocol 定义请求和响应的数据结构
// request.go 包含所有 API 请求的结构体定义
package protocol

import (
	"gopost/src/model"
)

// ==================== 文章相关请求 ====================

// REQNewArticle 创建新文章的请求结构
// binding 标签用于 Gin 的请求验证
type REQNewArticle struct {
	Title    string      `json:"title" binding:"required"`   // 文章标题，必需
	Tags     []string    `json:"tags"`                       // 文章标签列表
	Context  string      `json:"context" binding:"required"` // 文章内容，必需
	SNS      string      `json:"sns"`                        // 社交分享信息
	CurrUser *model.User // 当前登录用户（由中间件注入）
}

// REQUpdateArticle 更新文章的请求结构
// 嵌入 REQNewArticle 以复用字段验证
type REQUpdateArticle struct {
	REQNewArticle
	OldTitle string `json:"old_title" form:"old_title"` // 原文章标题，用于定位要更新的文章
}

// REQGetArticles 获取文章列表的请求结构
// 支持两种分页方式：
//   - 游标分页: 使用 time + size (原有方式)
//   - 页码分页: 使用 page + page_size (新增)
type REQGetArticles struct {
	Time     string `json:"time" form:"time"`           // 时间游标，用于获取比此时间更早的文章
	Size     string `json:"size" form:"size"`           // 每页数量 (游标分页)
	Page     string `json:"page" form:"page"`           // 页码，从 1 开始 (页码分页)
	PageSize string `json:"page_size" form:"page_size"` // 每页数量 (页码分页)，默认 10
}

// REQSearchArticle 搜索文章的请求结构
type REQSearchArticle struct {
	Keyword  string `json:"keyword" form:"keyword" binding:"required"` // 搜索关键词，必需
	Page     string `json:"page" form:"page"`                          // 页码，从 1 开始
	PageSize string `json:"page_size" form:"page_size"`                // 每页数量，默认 10
}

// REQGetArticlesByTag 按标签获取文章的请求结构
type REQGetArticlesByTag struct {
	Tag      string `json:"tag" form:"tag" binding:"required"` // 标签名称，必需
	Page     string `json:"page" form:"page"`                  // 页码，从 1 开始
	PageSize string `json:"page_size" form:"page_size"`        // 每页数量，默认 10
}

// REQGetArticle 获取单篇文章的请求结构
type REQGetArticle struct {
	Title    string      `form:"title" url:"title"` // 文章标题
	Aid      string      `form:"aid"`               // 文章 ID
	CurrUser *model.User // 当前登录用户
}

// REQDelArticle 删除文章的请求结构
type REQDelArticle struct {
	Title    string      `form:"title" url:"title"` // 要删除的文章标题
	CurrUser *model.User // 当前登录用户
}

// REQGetTags 获取文章标签的请求结构
type REQGetTags struct {
	Title string `form:"title" url:"title"` // 文章标题
}

// REQGetStat 获取统计信息的请求结构
type REQGetStat struct {
	Date string `form:"date"` // 日期筛选
}

// ==================== 评论相关请求 ====================

// REQNewReplay 创建新评论的请求结构
type REQNewReplay struct {
	Aid       string      `json:"aid" form:"aid"`                            // 文章 ID
	Title     string      `json:"title" form:"title"`                        // 文章标题
	Context   string      `json:"context" form:"context" binding:"required"` // 评论内容，必需
	CurrUser  *model.User // 当前登录用户
	IpAddress string      // 用户 IP 地址
}

// REQGetReplays 获取评论列表的请求结构
type REQGetReplays struct {
	Title    string      `form:"title" url:"title"` // 文章标题
	Aid      string      `form:"aid"`               // 文章 ID
	CurrUser *model.User // 当前登录用户
}

// REQDelReplays 删除评论的请求结构
type REQDelReplays struct {
	Rid      string      // 评论 ID
	CurrUser *model.User // 当前登录用户
}

// ==================== 用户相关请求 ====================

// REQGetUserInfo 获取用户信息的请求结构
type REQGetUserInfo struct {
	Username string      `form:"username" url:"username"` // 用户名
	Uid      string      `form:"uid"`                     // 用户 ID
	CurrUser *model.User // 当前登录用户
}

// REQSignin 登录请求结构
type REQSignin struct {
	Username string `json:"username" form:"username" binding:"required"` // 用户名，必需
	Password string `json:"password" form:"password" binding:"required"` // 密码，必需
	IP       string // 登录 IP
}

// REQSignUp 注册请求结构
// 嵌入 REQSignin 以复用字段验证
type REQSignUp struct {
	REQSignin
	Email string `json:"email" form:"email" binding:"required,email"` // 邮箱，必需且格式验证
}

// ==================== 心情相关请求 ====================

// REQNewMood 创建新心情的请求结构
type REQNewMood struct {
	Context string `form:"context" json:"context" binding:"required"` // 心情内容，必需
}

// REQLastMood 获取最新心情的请求结构
type REQLastMood struct {
	Limit int `form:"limit"` // 获取数量限制
}
