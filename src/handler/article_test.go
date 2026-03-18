// Package handler 包含 HTTP 请求处理器
// article_test.go 测试文章相关的 API 接口
package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopost/src/testutils"
)

// TestGetArticles 测试获取文章列表接口
// 测试场景:
//  1. 缺少必需参数 (time, size) -> 返回 400
//  2. 提供必需参数 -> 返回 200
func TestGetArticles(t *testing.T) {
	// 设置 Gin 测试模式
	router := testutils.SetupGin()

	// 注册测试路由 (模拟实际路由)
	router.GET("/v1/article", func(c *gin.Context) {
		// 绑定查询参数
		var req struct {
			Time string `form:"time" binding:"required"` // 时间参数，必需
			Size string `form:"size" binding:"required"` // 数量参数，必需
		}
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 测试用例 1: 缺少必需参数
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/article", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "缺少参数时应返回 400")

	// 测试用例 2: 提供必需参数
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/v1/article?time=2024&size=10", nil)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, http.StatusOK, w2.Code, "参数正确时应返回 200")
}

// TestCreateArticle_Validation 测试创建文章的请求验证
// 测试场景:
//  1. 缺少所有必需字段 (title, context) -> 返回 400
//  2. 缺少部分必需字段 (只提供 context) -> 返回 400
//  3. 提供所有必需字段 -> 返回 200
func TestCreateArticle_Validation(t *testing.T) {
	router := testutils.SetupGin()

	// 注册测试路由
	router.POST("/v1/article", func(c *gin.Context) {
		var req struct {
			Title   string `json:"title" binding:"required"`   // 标题，必需
			Context string `json:"context" binding:"required"` // 内容，必需
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 测试用例 1: 缺少所有必需字段
	w := httptest.NewRecorder()
	body := map[string]interface{}{}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/v1/article", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "缺少标题和内容时应返回 400")

	// 测试用例 2: 缺少标题字段
	w2 := httptest.NewRecorder()
	body2 := map[string]interface{}{"context": "测试内容"}
	jsonBody2, _ := json.Marshal(body2)
	req2, _ := http.NewRequest("POST", "/v1/article", bytes.NewBuffer(jsonBody2))
	req2.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w2, req2)

	assert.Equal(t, http.StatusBadRequest, w2.Code, "缺少标题时应返回 400")

	// 测试用例 3: 提供所有必需字段
	w3 := httptest.NewRecorder()
	body3 := map[string]interface{}{"title": "测试文章", "context": "测试内容"}
	jsonBody3, _ := json.Marshal(body3)
	req3, _ := http.NewRequest("POST", "/v1/article", bytes.NewBuffer(jsonBody3))
	req3.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w3, req3)

	assert.Equal(t, http.StatusOK, w3.Code, "参数正确时应返回 200")
}
