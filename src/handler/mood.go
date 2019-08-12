package handler

import (
	"github.com/gin-gonic/gin"
	"gopost/src/protocol"
	"gopost/src/service"
)

func NewMood(c *gin.Context) {
	var (
		req protocol.REQNewMood
		err error
	)

	defer func() {
		doResponse(c, err)
	}()

	if err = c.Bind(&req); err != nil {
		return
	}
	if err = service.NewMood(&req); err != nil {
		return
	}
}

func LastMood(c *gin.Context) {
	var (
		req protocol.REQLastMood
		res interface{}
		err error
	)
	defer func() {
		doResponse(c, err, res)
	}()
	if err = c.Bind(&req); err != nil {
		return
	}
	if req.Limit == 0 {
		req.Limit = 10
	}
	if res, err = service.LastMood(&req); err != nil {
		return
	}
}
