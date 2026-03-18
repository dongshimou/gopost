// Package logger 提供结构化日志功能
// 使用 logrus 库实现，支持日志级别和格式化输出
package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// 全局变量
var (
	// DEBUG 控制是否输出调试日志
	// 在 config.json 中配置，通过 SetDebugStatus() 设置
	DEBUG = false

	// Log 是全局 logrus 日志实例
	// 通过 InitLogger() 初始化
	Log = logrus.New()
)

// SetDebug 开启调试模式（硬编码方式）
// 建议使用 SetDebugStatus() 从配置文件读取
func SetDebug() {
	openDebug()
}

// openDebug 内部函数：开启调试模式
func openDebug() {
	DEBUG = true
}

// closeDebug 内部函数：关闭调试模式
// 注意：原代码有 bug，错误地设置为 true，现已修复
func closeDebug() {
	DEBUG = false
}

// SetDebugStatus 根据配置设置调试模式
// 参数 s: 来自 config.json 的 debug 字段
func SetDebugStatus(s bool) {
	if Log != nil {
		Log.WithField("status", s).Info("DEBUG status changed")
	}
	DEBUG = s
}

// Debug 输出调试级别日志
// 仅在 DEBUG=true 时输出
func Debug(args ...interface{}) {
	if !DEBUG {
		return
	}
	if Log != nil {
		Log.Debug(args...)
	}
}

// Error 输出错误级别日志
// 始终输出，不受 DEBUG 开关影响
// TODO: 可扩展为发送错误警报
func Error(args ...interface{}) {
	if Log != nil {
		Log.Error(args...)
	}
}

// Print 输出信息级别日志
func Print(args ...interface{}) {
	if Log != nil {
		Log.Info(args...)
	}
}

// InitLogger 初始化 logrus 日志系统
// 参数 debug: 是否开启调试级别日志
// 配置:
//   - 输出到 stdout
//   - debug=true 时输出 Debug 级别，否则输出 Info 级别
//   - 使用文本格式，包含完整时间戳
func InitLogger(debug bool) {
	Log.SetOutput(os.Stdout)

	if debug {
		Log.SetLevel(logrus.DebugLevel)
	} else {
		Log.SetLevel(logrus.InfoLevel)
	}

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
