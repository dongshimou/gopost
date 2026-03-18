// Package testutils 提供测试辅助工具
// 包含 HTTP 请求构建、响应解析、断言封装等功能
package testutils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// SetupGin 创建用于测试的 Gin 引擎
// 自动设置为测试模式 (gin.TestMode)
func SetupGin() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

// CreateTestRequest 创建 HTTP 测试请求
// 参数:
//   - method: HTTP 方法 (GET, POST, PUT, DELETE 等)
//   - path: 请求路径
//   - body: 请求体 (会被序列化为 JSON)
func CreateTestRequest(method, path string, body interface{}) *http.Request {
	var reqBody []byte
	if body != nil {
		reqBody, _ = json.Marshal(body)
	}
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	return req
}

// CreateJSONRequest 创建 JSON POST 请求
// 是 CreateTestRequest 的快捷方式，固定为 POST 方法
func CreateJSONRequest(path string, body interface{}) *http.Request {
	return CreateTestRequest("POST", path, body)
}

// GetResponseBody 从测试响应中提取 JSON body
// 返回 map[string]interface{} 格式的响应数据
func GetResponseBody(w *httptest.ResponseRecorder) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &result)
	return result
}

// AssertEqual 断言两个值相等
// 封装 testify 的 assert.Equal
func AssertEqual(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) {
	assert.Equal(t, expected, actual, msgAndArgs...)
}

// AssertNoError 断言错误为 nil
func AssertNoError(t assert.TestingT, err error, msgAndArgs ...interface{}) {
	assert.Nil(t, err, msgAndArgs...)
}

// AssertTrue 断言值为 true
func AssertTrue(t assert.TestingT, value bool, msgAndArgs ...interface{}) {
	assert.True(t, value, msgAndArgs...)
}

// AssertFalse 断言值为 false
func AssertFalse(t assert.TestingT, value bool, msgAndArgs ...interface{}) {
	assert.False(t, value, msgAndArgs...)
}

// AssertContains 断言字符串包含子串
func AssertContains(t assert.TestingT, s, contains string, msgAndArgs ...interface{}) {
	assert.Contains(t, s, contains, msgAndArgs...)
}
