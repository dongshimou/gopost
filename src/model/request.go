package model

type REQNewArticle struct {
	Title    string   `json:"title" binding:"required"`
	Tags     []string `json:"tags"`
	Context  string   `json:"context" binding:"required"`
	CurrUser *User    `json:"curr_user" form:"curr_user"`
}
type REQGetArticle struct {
	Title    string `form:"title" url:"title"`
	Aid      string `form:"aid"`
	CurrUser *User  `form:"curr_user"`
}
type REQNewReplay struct {
	Aid      string `json:"aid" form:"aid"`
	Title    string `json:"title" form:"title"`
	Context  string `json:"context" form:"context"`
	CurrUser *User  `form:"curr_user"`
}
type REQGetReplays struct {
	Title    string `form:"title" url:"title"`
	Aid      string `form:"aid" `
	CurrUser *User  `form:"curr_user"`
}
type REQGetUserInfo struct {
	Username string `form:"username" url:"username"`
	Uid      string `form:"uid"`
	CurrUser *User  `form:"curr_user"`
}

type REQSignin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	IP       string `json:"ip" form:"ip"`
}

type REQSignUp struct {
	REQSignin
	Email string `json:"email"`
}
