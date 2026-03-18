# Changelog

All notable changes to this project will be documented in this file.

## [0.1] - 2026-03-18

### Added
- 添加文章标题搜索接口 `GET /v1/article/search?keyword=xxx&size=10`
- 添加按标签获取文章接口 `GET /v1/article/tag?tag=xxx&size=10`
- 集成 `stretchr/testify` 测试框架
- 添加测试工具函数 `src/testutils/testutils.go`
- 添加示例测试文件 `src/handler/article_test.go`
- 集成 `sirupsen/logrus` 结构化日志系统
- 添加 `src/logger/common.go` 日志模块
- 添加中文注释到核心文件

### Fixed
- 修复 `logger.SetDebug()` 的 bug (closeDebug 错误地设置 DEBUG=true)
- 修复 `src/handler/article.go` 中的语法错误
- 修复 `main.go` 正确加载 `config.json` 中的 debug 配置
- 修复 logger 函数的空指针检查

### Changed
- 将 `c.Bind()` 改为 `c.ShouldBindJSON()` 以支持显式验证
- 为请求结构体添加 `binding` 验证标签
- 移除未使用的 `custom_validation.go`
