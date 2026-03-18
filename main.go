// Package main 是 gopost 项目的入口包
// gopost 是一个基于 Gin 框架的 RESTful 博客系统
package main

import (
	"gopost/src/base"
	"gopost/src/logger"
	v1 "gopost/src/v1"
)

// main 是程序入口
// 执行流程:
// 1. 初始化日志系统 (根据 config.json 的 debug 配置)
// 2. 初始化 V1 版本 API (数据库连接、路由注册)
// 3. 注册路由并启动 HTTP 服务
func main() {
	// 初始化日志系统，使用配置文件中的 debug 值
	// config.json 已在 base.init() 中自动加载
	logger.InitLogger(logger.DEBUG)

	// 初始化 V1 API：连接数据库、创建默认用户
	v1.InitV1()

	// 注册路由
	base.AddRoutes(v1.GetRoutes())

	// 启动 HTTP 服务
	base.StartService()
}
