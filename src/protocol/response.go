package protocol

import (
	"model"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type RESSignIn struct {
	Uid            uint   `json:"uid"`
	Username       string `json:"username"`
	Token          string `json:"token"`
	Permission     int    `json:"permission"`
	CreateDatetime string `json:"create_datetime"`
	UpdateDatetime string `json:"update_datetime"`
}
type RESSignUp struct {
}
type RESGetArticles struct {
	Articles []RESGetArticle `json:"articles"`
}
type RESGetArticle struct {
	Aid            uint     `json:"aid"`
	Title          string   `json:"title"`
	Author         string   `json:"author"`
	Tags           []string `json:"tags"`
	Context        string   `json:"context"`
	CreateDatetime string   `json:"create_datetime"`
	EditDatetime   string   `json:"edit_datetime"`

	Next string `json:"next"`
	Prev string `json:"prev"`
}
type RESGetReplaysSingle struct {
	Rid            uint   `json:"rid"`
	Username       string `json:"username"`
	IpAddress      string `json:"ip_address"`
	Context        string `json:"context"`
	CreateDatetime string `json:"create_datetime"`
}
type RESGetReplays struct {
	Aid          uint                  `json:"aid"`
	ArticleTitle string                `json:"article_title"`
	Replays      []RESGetReplaysSingle `json:"replays"`
}

type RESGetUserInfoArticle struct {
	Title          string `json:"title"`
	CreateDatetime string `json:"create_datetime"`
}
type RESGetUserInfoReplay struct {
	Title          string `json:"title"`
	Replay         string `json:"replay"`
	CreateDatetime string `json:"create_datetime"`
}
type RESGetUserInfo struct {
	Uid            uint                    `json:"uid"`
	Username       string                  `json:"username"`
	Email          string                  `json:"email"`
	Permission     int                     `json:"permission"`
	CreateDatetime string                  `json:"create_datetime"`
	UpdateDatetime string                  `json:"update_datetime"`
	PostArticle    []RESGetUserInfoArticle `json:"post_article"`
	PostReplay     []RESGetUserInfoReplay  `json:"post_replay"`
}

type RESGetTags struct {
	Names []string `json:"names"`
}
type ResGetStatSingle struct {
	Date string `json:"date"`
	Count int `json:"count"`
}
type RESGetStat struct {
	List []ResGetStatSingle `json:"list"`
}

type RESLastMood struct {
	List []model.Mood `json:"list"`
}