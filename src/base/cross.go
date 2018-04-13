package base

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"log"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOriginFunc = func(origin string) bool {
		log.Println("输入域名:", origin)
		return true
	}
	config.AllowCredentials = true
	config.AllowOrigins = []string{"*",}

	return cors.New(config)
}