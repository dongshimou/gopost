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

const (
	pathVer = "/v1"
)

func InitV1() error {
	err := controller.InitDB()
	if err != nil {
		return err
	}
	scfg := GetConfig().Server
	if logger.DEBUG {
		controller.GetDB().Create(&model.User{
			Name:     scfg.Name,
			Password: utility.EncryptPassword(scfg.Pass),
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
		user.Name = GetConfig().Server.Name
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
		"GetAllArticle",
		"GET",
		pathVer + "/article",
		handler.GetArticles,
		MakeAuth(model.Article_Read),
	},
	Route{
		"CreateArticle",
		"POST",
		pathVer + "/article",
		handler.CreateArticle,
		MakeAuth(model.Article_Create),
	},
	Route{
		"UpdateArticle",
		"POST",
		pathVer + "/article/update/:title",
		handler.UpdateArticle,
		MakeAuth(model.Article_Update),
	},
	Route{
		"GetArticle",
		"GET",
		pathVer + "/article/:title",
		handler.GetArticle,
		MakeAuth(model.Article_Read),
	},
	Route{
		"DelArticle",
		"DELETE",
		pathVer + "/article/:title",
		handler.DelArticle,
		MakeAuth(model.Article_Delete),
	},
	Route{
		"GetTags",
		"GET",
		pathVer + "/tags/:title",
		handler.GetTags,
		MakeAuth(model.Article_Read),
	},
	Route{
		"GetAllTags",
		"GET",
		pathVer + "/tags",
		handler.GetAllTags,
		MakeAuth(model.Article_Read),
	},
	Route{
		"PostReplay",
		"POST",
		pathVer + "/replay",
		handler.PostReplay,
		MakeAuth(model.Replay_Create),
	},
	Route{
		"GetReplay",
		"GET",
		pathVer + "/replay/:title",
		handler.GetReplays,
		MakeAuth(model.Replay_Read),
	},
	Route{
		"DelReplay",
		"DELETE",
		pathVer + "/replay/:title/:count",
		handler.DelReplays,
		MakeAuth(model.Replay_Delete),
	},
	Route{
		"GetUserInfo",
		"GET",
		pathVer + "/user/:username",
		handler.GetUserInfo,
		MakeAuth(model.User_Read),
	},

	//不需要权限验证
	Route{
		"SignIn",
		"POST",
		pathVer + "/sign/in",
		handler.SignIn,
		nil,
	},
	Route{
		"SignOut",
		"GET",
		pathVer + "/sign/out",
		handler.SignOut,
		nil,
	},
	Route{
		"SignUp",
		"POST",
		pathVer + "/sign/up",
		handler.SignUp,
		nil,
	},
	Route{
		"SignVerify",
		"GET",
		pathVer + "/sign/verify",
		handler.SignVerify,
		nil,
	},

	//rss
	Route{
		"RSS",
		"GET",
		pathVer + "/rss",
		handler.Rss,
		nil,
	},
}
