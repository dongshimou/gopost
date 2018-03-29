package v1

import (
	. "base"
	"controller"
	"handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitV1() error {
	err := controller.InitDB()
	if err != nil {
		return err
	}
	return nil
}
func GetRoutes() []Route {
	return routes
}
func Fail(c *gin.Context) {
	//c.String(http.StatusOK, "admin auth failed")
	c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"code": 1004,
			"msg":  "权限错误",
			"data": nil,
		})
}

var routes = []Route{
	Route{
		"NewPost",
		"POST",
		"/v1/new",
		handler.NewPost,
		nil,
	},
	Route{
		"Login",
		"POST",
		"/v1/login",
		handler.Login,
		nil,
	},
}
