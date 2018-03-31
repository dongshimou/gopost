package model

type REQNewArticle struct {
	Title   string   `json:"title" binding:"required"`
	Tag     []string `json:"tag"`
	Context string   `json:"context" binding:"required"`
}
type REQGetArticle struct {
	Title string `json:"title" form:"title" url:"title"`
	ID    string `json:"id" form:"title"`
}
type REQNewReplay struct {
	Aid string `json:"aid" form:"aid"`
	Context string `json:"context" form:"context"`
}
type REQLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type REQSignUp struct {
	REQLogin
	Email string `json:"email"`
}
