package model

import "encoding/xml"

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
type RESGetArticle struct {
	Aid            uint     `json:"aid"`
	Title          string   `json:"title"`
	Author         string   `json:"author"`
	Tags           []string `json:"tags"`
	Context        string   `json:"context"`
	ReplayCount    uint     `json:"replay_count"`
	CreateDatetime string   `json:"create_datetime"`
	EditDatetime   string   `json:"edit_datetime"`
}

type RESGetReplays struct {
	Aid          uint   `json:"aid"`
	ArticleTitle string `json:"article_title"`
	Replays      []struct {
		Rid            uint   `json:"rid"`
		Username       string `json:"username"`
		Context        string `json:"context"`
		CreateDatetime string `json:"create_datetime"`
	} `json:"replays"`
}

type RESGetUserInfo struct {
	Uid            uint   `json:"uid"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Permission     int    `json:"permission"`
	CreateDatetime string `json:"create_datetime"`
	UpdateDatetime string `json:"update_datetime"`
	PostArticle    []struct {
		Title          string `json:"title"`
		CreateDatetime string `json:"create_datetime"`
	} `json:"post_article"`
	PostReplay []struct {
		Title          string `json:"title"`
		Replay         string `json:"replay"`
		CreateDatetime string `json:"create_datetime"`
	} `json:"post_replay"`
}

type RssFeed struct {
	//有且仅有一个 channel
	XMLName xml.Name `xml:"rss"`
	Version string `xml:"version,attr"`
	Channel *RssChanel `xml:"channel"`
}
type RssChanel struct {
	XMLName xml.Name `xml:"channel"`
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Language    string    `xml:"language,omitempty"`
	Copyright   string    `xml:"copyright,omitempty"`
	PubDate     string    `xml:"pubDate,omitempty"`
	ManagingEditor string `xml:"managingEditor,omitempty"`
	LastBuildDate string `xml:"lastBuildDate,omitempty"`
	Item        []RssItem `xml:"item"`
}
type RssItem struct {
	XMLName xml.Name `xml:"item"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Author      string `xml:"author"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}
