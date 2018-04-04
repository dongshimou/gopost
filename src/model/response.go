package model

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
