package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service"
)

func Rss(c *gin.Context){
	xml,err:=service.Rss()
	if err!=nil{

	}
	c.XML(http.StatusOK,xml)
}
