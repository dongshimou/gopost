package main

import (
	"v1"
	"base"
)

func main() {
	v1.InitV1()
	base.AddRoutes(v1.GetRoutes())
	base.StartService()
}
