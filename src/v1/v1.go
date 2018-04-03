package v1

import (
	. "base"
	"controller"
	"github.com/gin-gonic/gin"
	"handler"
	"model"
	"net/http"
	"utility"
)

func InitV1() error {
	err := controller.InitDB()
	if err != nil {
		return err
	}

	//for test
	controller.GetDB().Create(&model.User{Name: "root", Password: "123456", Permission: 0x1111})

	return nil
}
func GetRoutes() []Route {
	return routes
}
func Fail(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"code": utility.ERROR_AUTH_CODE,
			"msg":  utility.ERROR_AUTH_MSG,
			"data": nil,
		})
}

var routes = []Route{
	Route{
		"PostNewArticle",
		"POST",
		"/v1/article",
		handler.PostNewArticle,
		nil,
	},
	Route{
		"Signin",
		"POST",
		"/v1/sign/in",
		handler.Signin,
		nil,
	},
	Route{
		"SignOut",
		"GET",
		"/v1/sign/out",
		handler.SignOut,
		nil,
	},
	Route{
		"SignUp",
		"POST",
		"/v1/sign/up",
		handler.SignUp,
		nil,
	},
	Route{
		"GetArticle",
		"GET",
		"/v1/article/:title",
		handler.GetArticle,
		nil,
	},
	Route{
		"PostNewReplay",
		"POST",
		"/v1/replay",
		handler.PostNewReplay,
		nil,
	},
	Route{
		"GetReplay",
		"GET",
		"/v1/replay/:title",
		handler.GetReplays,
		nil,
	},
	Route{
		"GetUserInfo",
		"GET",
		"/v1/user/:username",
		handler.GetUserInfo,
		nil,
	},
}
