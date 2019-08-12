package handler

import (
	"github.com/gin-gonic/gin"
	"gopost/src/service"
	"net/http"
)

func Rss(c *gin.Context) {
	xml, err := service.Rss()
	if err != nil {
		c.XML(http.StatusInternalServerError, err.Error())
	}
	c.XML(http.StatusOK, xml)
}
