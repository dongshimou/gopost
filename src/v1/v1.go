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
		controller.GetDB().Create(&model.User{
			Name:     "root",
			Password: "123456",
			Permission: utility.CreatePermission(
				model.Article_Read,
				model.Article_Update,
				model.Article_Delete,
				model.Article_Create,
				model.Replay_Read,
				model.Replay_Update,
				model.Replay_Delete,
				model.Replay_Create,
				model.User_Read,
				model.User_Update,
				model.User_Delete,
				model.User_Create,
			)})
	}

	return nil
}

func GetRoutes() []Route {
	return routes
}
func MakeAuth(prems ...int) gin.HandlersChain {
	return MakeHandler(handler.AuthDecorator(GetUserFromToken, prems...))
}
func GetUserFromToken(token string) (*model.User, error) {
	user := &model.User{}
	var err error
	if logger.DEBUG {
		user.Name = "root"
	} else {
		if user.Name, err = utility.ParseToken(token); err != nil {
			return nil, err
		}
	}
	db := controller.GetDB()
	if err = db.Model(user).Where(user).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
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
		"DelArticle",
		"DELETE",
		"/v1/article/:title",
		handler.DelArticle,
		MakeAuth(model.Article_Delete),
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
		"DelReplay",
		"DELETE",
		"/v1/replay/:title/:rid",
		handler.DelReplays,
		MakeAuth(model.Replay_Delete),
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
	Route{
		"SignVerify",
		"GET",
		"/v1/sign/verify",
		handler.SignVerify,
		nil,
	},
}
