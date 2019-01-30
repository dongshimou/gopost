package service

import (
	"model"
	"orm"
	"protocol"
)

func NewMood(req *protocol.REQNewMood)(error){
	db:=orm.Get()
	return db.Create(&model.Mood{Context:req.Context}).Error
}

func LastMood(req *protocol.REQLastMood)(res *protocol.RESLastMood,err error){
	db:=orm.Get()

	moods:=[]model.Mood{}
	if err:=db.Model(&model.Mood{}).Order("created_at desc").Limit(req.Limit).Find(&moods).Error;err!=nil{
		return nil,err
	}
	res=&protocol.RESLastMood{}
	res.List=moods
	return res,nil
}