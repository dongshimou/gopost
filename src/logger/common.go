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
