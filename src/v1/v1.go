package v1

import (
	. "base"
	"controller"
	"github.com/gin-gonic/gin"
	"handler"
	"logger"
	"model"
	"utility"
)

func InitV1() error {
	err := controller.InitDB()
	if err != nil {
		return err
	}

	if logger.DEBUG {
		controller.GetDB().Create(&model.User{Name: "root", Password: "123456", Permission: 0x1111})
	}

	return nil
}

func GetRoutes() []Route {
	return routes
}
func MakeAuth(prems ...int) gin.HandlersChain {
	return MakeHandler(handler.AuthDecorator(GetUserFromToken, Fail, prems...))
}
func Fail(c *gin.Context) {
	handler.DoResponseFail(c, utility.NewError(utility.ERROR_AUTH_CODE, utility.ERROR_AUTH_MSG))
}
func GetUserFromToken(token string) *model.User {
	user := &model.User{}

	if logger.DEBUG {
		user.Name = "root"
		db := controller.GetDB()
		if err := db.Model(user).Where(user).First(user).Error; err != nil {
			return nil
		}
	}

	return user
}

var routes = []Route{
	Route{
		"PostArticle",
		"POST",
		"/v1/article",
		handler.PostArticle,
		MakeAuth(model.Article_Create),
	},
	Route{
		"GetArticle",
		"GET",
		"/v1/article/:title",
		handler.GetArticle,
		MakeAuth(model.Article_Read),
	},
	Route{
		"PostReplay",
		"POST",
		"/v1/replay",
		handler.PostReplay,
		MakeAuth(model.Replay_Create),
	},
	Route{
		"GetReplay",
		"GET",
		"/v1/replay/:title",
		handler.GetReplays,
		MakeAuth(model.Replay_Read),
	},
	Route{
		"GetUserInfo",
		"GET",
		"/v1/user/:username",
		handler.GetUserInfo,
		MakeAuth(model.User_Read),
	},

	//不需要权限验证
	Route{
		"SignIn",
		"POST",
		"/v1/sign/in",
		handler.SignIn,
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
}
