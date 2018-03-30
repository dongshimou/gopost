package service

import (
	"controller"
	"model"
	"utility"
)

func isNullOrEmpty(v interface{}) bool {
	return utility.IsNullorEmpty(v)
}
func parse2uint(s string) (uint, error) {
	return utility.Parse2Uint(s)
}
func NewPost(req *model.REQNewPost) error {

	user := model.User{
		Name: "root",
	}
	tags := []model.Tag{}
	for i, _ := range req.Tag {
		tags = append(tags, model.Tag{
			Name: req.Tag[i],
		})
	}

	db := controller.GetDB().Where(&user).First(&user)
	if db.Error != nil {
		return db.Error
	}

	post := model.Post{
		Title:        req.Title,
		Context:      req.Context,
		Tags:         tags,
		AuthorID:     user.ID,
		EditPersonID: user.ID,
	}

	tx := db.Begin()

	tx = tx.Create(&post)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()
	return nil
}
func GetPost(req *model.REQGetPost) (*model.RESGetPost, error) {
	post := model.Post{}
	if !isNullOrEmpty(req.ID) {
		id, err := parse2uint(req.ID)
		if err != nil {
			return nil, err
		}
		post.ID = id
		goto query
	}

	if !isNullOrEmpty(req.Title) {

		post.Title = req.Title
		goto query
	}

	return nil, utility.NewError(utility.ERROR_REQUEST_CODE, utility.ERROR_REQUEST_MSG)

query:
	db := controller.GetDB()
	db = db.Where(&post).First(&post).Model(&post)
	db.Related(&post.Author)
	db.Related(&post.Tags)
	if db.Error != nil {
		return nil, db.Error
	}
	res := model.RESGetPost{
		Title:  post.Title,
		Author: post.Author.Name,
		Tags: func(tags []model.Tag) []string {
			ts := []string{}
			for i, _ := range tags {
				ts = append(ts, tags[i].Name)
			}
			return ts
		}(post.Tags),
		Context:        post.Context,
		CreateDatetime: post.CreatedAt,
		EditDatetime:   post.UpdatedAt,
	}
	return &res, nil
}
func Login(req *model.REQLogin) (*model.RESLogin, error) {

	return nil, nil
}
