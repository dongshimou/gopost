package main

import (
	"base"
	"logger"
	"v1"
)

func main() {

	logger.SetDebug()

	v1.InitV1()
	base.AddRoutes(v1.GetRoutes())
	base.StartService()
}
