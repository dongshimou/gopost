package base

import "github.com/gin-gonic/gin"

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
	router.Run(":" + GetConfig().Server.Port)
}
