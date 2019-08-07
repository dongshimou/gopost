package main

import (
	"gopost/src/base"
	"gopost/src/logger"
	"gopost/src/v1"
)

func main() {

	logger.SetDebug()

	v1.InitV1()
	base.AddRoutes(v1.GetRoutes())
	base.StartService()
}
