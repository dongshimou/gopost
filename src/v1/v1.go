package v1

import (
	. "base"
	"controller"
	"github.com/gin-gonic/gin"
	"handler"
	"model"
	"net/http"
)

func InitV1() error {
	err := controller.InitDB()
	if err != nil {
		return err
	}
	controller.GetDB().Create(&model.User{Name: "root", Password: "123456"})
	//controller.GetDB().Model(&model.Post{}).AddForeignKey("author_id","users(id)","RESTRICT","RESTRICT")
	return nil
}
func GetRoutes() []Route {
	return routes
}
func Fail(c *gin.Context) {
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
	Route{
		"Logout",
		"GET",
		"/v1/logout",
		handler.Logout,
		nil,
	},
	Route{
		"SignUp",
		"POST",
		"/v1/signup",
		handler.SignUp,
		nil,
	},
	Route{
		"GetArticle",
		"GET",
		"/v1/post/:title",
		handler.GetArticle,
		nil,
	},
	Route{
		"NewReplay",
		"POST",
		"/v1/replay/:aid",
		handler.NewReplay,
		nil,
	},
}
