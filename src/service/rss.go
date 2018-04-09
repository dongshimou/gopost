package service

import (
	"model"
	"time"
	"controller"
	"base"
)

//github.com/gorilla/feeds
func Rss()(xml interface{},err error){
	scfg:=base.GetConfig().Server

	proto:="http://"
	host:="127.0.0.1"
	port:=scfg.Port
	address:=proto+host+":"+port
	apath:="/v1/article/"
	rsspath:="/v1/rss"
	res:=model.RssFeed{
		Version:"2.0",
		Channel:&model.RssChanel{
		Title:scfg.Name,
		Link:address+rsspath,
		Description:"go post",
		Language:"zh-cn",
		PubDate:time.Now().Format(time.RFC1123),
		//LastBuildDate:time.Now().Format(time.RFC1123),
		},
	}

	db:=controller.GetDB()
	list:=[]model.Article{}
	items:=[]model.RssItem{}

	db.Model(&model.Article{}).
	Select(buildArgs(",",model.Table_Article_Title,model.Table_Article_AuthorName,model.DB_created_at)).
	Find(&list)

	for i,_:=range list{
		v:=&list[i]
		items=append(items,model.RssItem{
			Title:v.Title,
			Link:address+apath+v.Title,
			Author:v.AuthorName,
			PubDate:v.CreatedAt.Format(time.RFC1123),
			Description:v.Title,
		})
	}
	res.Channel.Item=items
	if len(list)!=0{
		res.Channel.LastBuildDate=list[len(list)-1].CreatedAt.Format(time.RFC1123)
	}
	return res,nil
}
