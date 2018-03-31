package service

import (
	"utility"
	"controller"
	"model"
)

func NewPost(req *model.REQNewArticle) error {

	user := model.User{
		Name: "root",
	}
	db := controller.GetDB()

	tags := []*model.Tag{}
	for i, _ := range req.Tag {
		t:=model.Tag{Name: req.Tag[i],}
		db.FirstOrCreate(&t,&t)
		tags = append(tags,&t )
	}

	if db.Where(&user).First(&user).Error != nil {
		return db.Error
	}
	post := model.Post{
		Title:    req.Title,
		Context:  req.Context,
		Tags:     tags,
		AuthorID: user.ID,
		EditorID: user.ID,
	}

	tx := db.Begin()
	var err error
	if err = tx.Save(&post).Error; err != nil {
		goto rollback
	}
	goto commit

rollback:
	tx.Rollback()
	return err
commit:
	tx.Commit()
	return nil
}
func GetPost(req *model.REQGetArticle) (*model.RESGetArticle, error) {
	post := model.Post{}
	if !isNullOrEmpty(req.ID) {
		id, err := parseID(req.ID)
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
	//查询post
	db.Where(&post).First(&post)
	//查询tags
	db.Model(&post).Related(&post.Tags, "tags")
	//db.Model(&post).Association("tags").Find(&post.Tags)
	//查询user
	db.Model(&post).Related(&post.Author, "author_id")

	if db.Error != nil {
		return nil, db.Error
	}
	res := model.RESGetArticle{
		Aid:    post.ID,
		Title:  post.Title,
		Author: post.Author.Name,
		Tags: func(tags []*model.Tag) []string {
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

func NewReplay(req *model.REQNewReplay)(err error) {
	if isNullOrEmpty(req.Aid) || isNullOrEmpty(req.Context) {
return utility.NewError(utility.ERROR_REQUEST_CODE, utility.ERROR_REQUEST_MSG)
	}
	aid, err := parseID(req.Aid)
	if err != nil {
		return err
	}
	db:=controller.GetDB()
	user:=model.User{
		Name:"root",
	}
	if err=db.Model(&user).Where(&user).First(&user).Error;err!=nil{
		return err
	}
	replay:=model.Replay{
		PostID:aid,
		UserID:user.ID,
		Context:req.Context,
	}
	tx:=db.Begin()
	if err=tx.Save(&replay).Error;err!=nil{
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}