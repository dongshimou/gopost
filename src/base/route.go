package base

import (
	"github.com/gin-gonic/gin"
	"logger"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
	AuthHandler gin.HandlersChain
}

var routes []Route

func AddRoutes(r []Route) {
	routes = append(routes, r...)
}
func MakeHandler(args ...gin.HandlerFunc) gin.HandlersChain {
	res := gin.HandlersChain{}
	res = append(res, args...)
	return res
}
func StartService() {
	router := gin.Default()
	//jsonp 中间件
	router.Use(JsonP())
	//option 中间件
	router.Use(MyCors())

	for _, route := range routes {
		pattern := route.Pattern
		hanlders := gin.HandlersChain{}
		if route.AuthHandler != nil {
			hanlders = append(hanlders, route.AuthHandler...)
		}
		hanlders = append(hanlders, route.HandlerFunc)

		switch route.Method {
		case "GET":
			router.GET(pattern, hanlders...)
		case "POST":
			router.POST(pattern, hanlders...)
		case "DELETE":
			router.DELETE(pattern, hanlders...)
		}
	}
	address := ":" + GetConfig().Server.Port
	if GetConfig().Server.TLS {
		logger.Debug("cert file:", GetConfig().Server.CertFile)
		logger.Debug("key file:", GetConfig().Server.KeyFile)
		router.RunTLS(address, GetConfig().Server.CertFile, GetConfig().Server.KeyFile)
	} else {
		logger.Debug("no tls")
		router.Run(address)
	}
}
