package handler

import (
	"github.com/gin-gonic/gin"
	"protocol"
	"service"
)

func NewMood(c *gin.Context){
	var req protocol.REQNewMood
	var err error

	defer func() {
		if err!=nil{
			DoResponseFail(c,err)
		}else{
			DoResponseOK(c,nil)
		}
	}()

	if err=c.Bind(&req);err!=nil{
		return
	}
	if err=service.NewMood(&req);err!=nil{
		return
	}
}

func LastMood(c *gin.Context){
	var req protocol.REQLastMood
	var res *protocol.RESLastMood
	var err error
	defer func() {
		if err!=nil{
			DoResponseFail(c,err)
		}else{
			DoResponseOK(c,res)
		}
	}()
	if err=c.Bind(&req);err!=nil{
		return
	}
	if req.Limit==0{
		req.Limit=20
	}
	if res,err=service.LastMood(&req);err!=nil{
		return
	}
}