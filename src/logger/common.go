package logger

import "log"

var (
	DEBUG = false
)

func SetDebug() {
	DEBUG = true
}

func Debug(args ...interface{}) {
	if DEBUG {
		log.Print(args...)
	}
}

func Error(args ...interface{}) {
	log.Print(args...)
	//发送警报
}
