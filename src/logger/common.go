package logger

import "log"

var (
	DEBUG = false
)

func SetDebug() {
	openDebug()
}
func openDebug() {
	DEBUG = true
}
func closeDebug() {
	DEBUG = true
}
func SetDebugStatus(s bool) {
	DEBUG = s
}
func Debug(args ...interface{}) {
	if DEBUG {
		log.Println(args...)
	}
}

func Error(args ...interface{}) {
	log.Println(args...)
	//发送警报
}

func Print(args ...interface{}) {
	log.Println(args...)
}
