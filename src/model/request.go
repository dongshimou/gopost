package model

type REQNewArticle struct {
	Title   string   `json:"title" binding:"required"`
	Tags     []string `json:"tags"`
	Context string   `json:"context" binding:"required"`
	CurrUser *REQCurrUser
}
type REQGetArticle struct {
	Title string `json:"title" form:"title" url:"title"`
	Aid   string `json:"aid" form:"aid"`
	CurrUser *REQCurrUser
}
type REQNewReplay struct {
	Aid     string `json:"aid" form:"aid"`
	Title   string `json:"title" form:"title"`
	Context string `json:"context" form:"context"`
	CurrUser *REQCurrUser
}
type REQGetReplays struct {
	Title string `json:"title" form:"title" url:"title"`
	Aid   string `json:"aid" form:"aid" `
	CurrUser *REQCurrUser
}
type REQGetUserInfo struct {
	Username string `json:"username" form:"username" url:"username"`
	Uid      string `json:"uid" form:"uid"`
	CurrUser *REQCurrUser
}

type REQCurrUser struct {
	CurrToken string
	CurrUser string
	CurrID uint
	CurrPer int
	CurrEmail string
}

type REQLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	IP string `json:"ip" form:"ip"`
}

type REQSignUp struct {
	REQLogin
	Email string `json:"email"`
}

