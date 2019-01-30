package service

import (
	"base"
	"github.com/gorilla/feeds"
	"model"
	"orm"
	"time"
)

//github.com/gorilla/feeds
func Rss() (xml interface{}, err error) {
	scfg := base.GetConfig().Server

	apath := "/v1/article/"

	now := time.Now()
	feed := &feeds.Feed{
		Title:       scfg.Name,
		Link:        &feeds.Link{Href: scfg.Link},
		Description: scfg.Description,
		Author:      &feeds.Author{Name: scfg.Name, Email: scfg.Email},
		Created:     now,
	}

	db := orm.Get()

	articles := []model.Article{}

	//查找最后十个
	db.Model(&model.Article{}).
		Select(buildArgs(",", model.Table_Article_Title, model.DB_created_at, model.Table_Article_AuthorName)).
		Limit(10).
		Find(&articles)

	for i, _ := range articles {
		v := &articles[i]
		feed.Add(&feeds.Item{
			Title:       v.Title,
			Link:        &feeds.Link{Href: scfg.Link + apath + v.Title},
			Description: v.Title,
			Author:      &feeds.Author{Name: v.AuthorName},
			Created:     v.CreatedAt,
		})
	}

	//rss,err:=feed.ToRss()
	//if err!=nil{
	//	logger.Print(err)
	//}
	//<?xml version="1.0" encoding="UTF-8"?> len=38

	rss := feeds.Rss{feed}
	return rss.FeedXml(), err
}
