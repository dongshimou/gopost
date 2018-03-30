package model

type REQNewPost struct {
	Title   string   `json:"title" binding:"required"`
	Tag     []string `json:"tag"`
	Context string   `json:"context" binding:"required"`
}
type REQGetPost struct {
	Title string `json:"title" form:"title" url:"title"`
	ID    string `json:"id" form:"title"`
}

type REQLogin struct {
}
