package service

import (
	"model"
	"orm"
)

func NewMood(req *model.REQNewMood)(error){
	db:=orm.Get()
	return db.Create(&model.Mood{Context:req.Context}).Error
}

func LastMood(req *model.REQLastMood)(res *model.RESLastMood,err error){
	db:=orm.Get()

	moods:=[]model.Mood{}
	if err:=db.Model(&model.Mood{}).Order("created_at desc").Limit(req.Limit).Find(&moods).Error;err!=nil{
		return nil,err
	}
	res=&model.RESLastMood{}
	res.List=moods
	return res,nil
}