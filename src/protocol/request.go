package protocol

import (
	"gopost/src/model"
)

type REQNewArticle struct {
	Title    string   `json:"title" binding:"required"`
	Tags     []string `json:"tags"`
	Context  string   `json:"context" binding:"required"`
	SNS      string   `json:"sns"`
	CurrUser *model.User
}
type REQUpdateArticle struct {
	REQNewArticle
	OldTitle string
}
type REQGetArticles struct {
	Time string `json:"time" form:"time" binding:"required"`
	Size string `json:"size" form:"size" binding:"required"`
}
type REQGetArticle struct {
	Title    string `form:"title" url:"title"`
	Aid      string `form:"aid"`
	CurrUser *model.User
}
type REQDelArticle struct {
	Title    string `form:"title" url:"title"`
	CurrUser *model.User
}
type REQGetTags struct {
	Title string `form :"title" url:"title"`
}
type REQGetStat struct {
	Date string `form:"date"`
}
type REQNewReplay struct {
	Aid       string `json:"aid" form:"aid"`
	Title     string `json:"title" form:"title"`
	Context   string `json:"context" form:"context"`
	CurrUser  *model.User
	IpAddress string
}
type REQGetReplays struct {
	Title    string `form:"title" url:"title"`
	Aid      string `form:"aid" `
	CurrUser *model.User
}
type REQDelReplays struct {
	Rid      string
	CurrUser *model.User
}
type REQGetUserInfo struct {
	Username string `form:"username" url:"username"`
	Uid      string `form:"uid"`
	CurrUser *model.User
}

type REQSignin struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	IP       string
}

type REQSignUp struct {
	REQSignin
	Email string `json:"email" form:"email" binding:"required,email"`
}

type REQNewMood struct {
	Context string `form:"context" json:"context" binding:"required"`
}
type REQLastMood struct {
	Limit int `form:"limit"`
}
