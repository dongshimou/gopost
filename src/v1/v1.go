// Package v1 提供 API v1 版本的路由和初始化
// 包含路由注册、认证中间件、用户权限验证等功能
package v1

import (
	"github.com/gin-gonic/gin"
	. "gopost/src/base"
	"gopost/src/handler"
	"gopost/src/logger"
	"gopost/src/model"
	"gopost/src/orm"
	"gopost/src/utility"
)

// pathVer 是 API 版本前缀
const (
	pathVer = "/v1"
)

// InitV1 初始化 V1 版本 API
// 执行流程:
//  1. 初始化数据库连接
//  2. 在调试模式下创建默认管理员用户
func InitV1() error {
	// 初始化数据库连接和表结构
	err := orm.InitDB()
	if err != nil {
		return err
	}

	// 调试模式下创建默认用户
	scfg := GetConfig().Server
	if logger.DEBUG {
		orm.Get().Create(&model.User{
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

// GetRoutes 返回所有 V1 API 路由
func GetRoutes() []Route {
	return routes
}

// MakeAuth 创建认证中间件链
// 参数 prems: 需要的权限列表
// 返回: Gin 处理器链
func MakeAuth(prems ...int) gin.HandlersChain {
	return MakeHandler(handler.AuthDecorator(GetUserFromToken, prems...))
}

// GetUserFromToken 从 Token 获取用户信息
// 参数 token: JWT 认证令牌
// 返回: 用户模型和错误
// 在调试模式下，直接使用配置文件中的用户名
func GetUserFromToken(token string) (*model.User, error) {
	user := &model.User{}
	var err error

	// 调试模式：跳过 Token 验证
	if logger.DEBUG {
		user.Name = GetConfig().Server.Name
	} else {
		// 生产模式：解析 JWT Token
		if user.Name, err = utility.ParseToken(token); err != nil {
			return nil, err
		}
	}

	// 从数据库获取完整用户信息
	db := orm.Get()
	if err = db.Model(user).Where(user).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// routes 定义所有 API 路由
// 每个路由包含: 名称、HTTP方法、路径、处理器、认证中间件
var routes = []Route{
	// ==================== 统计接口 ====================
	Route{
		"GetIpStat", // 获取 IP 访问统计
		"GET",
		pathVer + "/stat",
		handler.GetStat,
		nil, // 无需认证
	},

	// ==================== 心情接口 ====================
	Route{
		"CreateMood", // 创建心情
		"GET,POST",
		pathVer + "/mood/new",
		handler.NewMood,
		nil,
	},
	Route{
		"GetMoodLast", // 获取最新心情
		"GET",
		pathVer + "/mood/last",
		handler.LastMood,
		nil,
	},

	// ==================== 文章接口 ====================
	Route{
		"GetAllArticle", // 获取文章列表
		"GET",
		pathVer + "/article",
		handler.GetArticles,
		append(MakeAuth(model.Article_Read), handler.StatIp), // 需要文章读取权限
	},
	Route{
		"SearchArticles", // 搜索文章
		"GET",
		pathVer + "/article/search",
		handler.SearchArticles,
		MakeAuth(model.Article_Read), // 需要文章读取权限
	},
	Route{
		"GetArticlesByTag", // 按标签获取文章
		"GET",
		pathVer + "/article/tag",
		handler.GetArticlesByTag,
		MakeAuth(model.Article_Read), // 需要文章读取权限
	},
	Route{
		"CreateArticle", // 创建文章
		"POST",
		pathVer + "/article",
		handler.CreateArticle,
		MakeAuth(model.Article_Create), // 需要文章创建权限
	},
	Route{
		"UpdateArticle", // 更新文章
		"PUT",
		pathVer + "/article/:oldtitle",
		handler.UpdateArticle,
		MakeAuth(model.Article_Update), // 需要文章更新权限
	},
	Route{
		"UpdateArticle", // 更新文章 (POST 方式)
		"POST",
		pathVer + "/article/update/:oldtitle",
		handler.UpdateArticle,
		MakeAuth(model.Article_Update),
	},
	Route{
		"GetArticle", // 获取单篇文章
		"GET",
		pathVer + "/article/:title",
		handler.GetArticle,
		append(MakeAuth(model.Article_Read), handler.StatIp),
	},
	Route{
		"DelArticle", // 删除文章
		"DELETE",
		pathVer + "/article/:title",
		handler.DelArticle,
		MakeAuth(model.Article_Delete), // 需要文章删除权限
	},
	Route{
		"GetTags", // 获取文章标签
		"GET",
		pathVer + "/tags/:title",
		handler.GetTags,
		MakeAuth(model.Article_Read),
	},
	Route{
		"GetAllTags", // 获取所有标签
		"GET",
		pathVer + "/tags",
		handler.GetAllTags,
		MakeAuth(model.Article_Read),
	},

	// ==================== 评论接口 ====================
	Route{
		"PostReplay", // 创建评论
		"POST",
		pathVer + "/replay",
		handler.PostReplay,
		MakeAuth(model.Replay_Create), // 需要评论创建权限
	},
	Route{
		"GetReplay", // 获取评论列表
		"GET",
		pathVer + "/replay/:title",
		handler.GetReplays,
		MakeAuth(model.Replay_Read),
	},
	Route{
		"DelReplay", // 删除评论
		"DELETE",
		pathVer + "/replay/:rid",
		handler.DelReplays,
		MakeAuth(model.Replay_Delete), // 需要评论删除权限
	},

	// ==================== 用户接口 ====================
	Route{
		"GetUserInfo", // 获取用户信息
		"GET",
		pathVer + "/user/:username",
		handler.GetUserInfo,
		MakeAuth(model.User_Read),
	},

	// ==================== 认证接口 (无需权限验证) ====================
	Route{
		"SignIn", // 登录
		"POST",
		pathVer + "/sign/in",
		handler.SignIn,
		nil, // 无需认证
	},
	Route{
		"SignOut", // 登出
		"GET",
		pathVer + "/sign/out",
		handler.SignOut,
		nil,
	},
	Route{
		"SignUp", // 注册
		"POST",
		pathVer + "/sign/up",
		handler.SignUp,
		nil,
	},
	Route{
		"SignVerify", // 验证登录状态
		"GET",
		pathVer + "/sign/verify",
		handler.SignVerify,
		nil,
	},

	// ==================== RSS 订阅接口 ====================
	Route{
		"RSS", // 获取 RSS 订阅
		"GET",
		pathVer + "/rss",
		handler.Rss,
		nil,
	},
}
