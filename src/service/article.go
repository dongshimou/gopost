package service

import (
	"controller"
	"model"
	"utility"
)

func PostArticle(req *model.REQNewArticle) error {

	user := req.CurrUser
	if user == nil {
		return utility.NewError(utility.ERROR_AUTH_CODE, utility.ERROR_MSG_UNKNOW_USER)
	}
	db := controller.GetDB()

	tags := []*model.Tag{}
	for i, _ := range req.Tags {
		t := model.Tag{Name: req.Tags[i]}
		db.FirstOrCreate(&t, &t)
		tags = append(tags, &t)
	}

	post := model.Article{
		Title:    req.Title,
		Context:  req.Context,
		Tags:     tags,
		AuthorName:user.Name,
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
func GetArticle(req *model.REQGetArticle) (*model.RESGetArticle, error) {
	article := model.Article{}
	if !isNullOrEmpty(req.Aid) {
		aid, err := parseID(req.Aid)
		if err != nil {
			return nil, err
		}
		article.ID = aid
		goto query
	}
	if !isNullOrEmpty(req.Title) {
		article.Title = req.Title
		goto query
	}
	return nil, utility.NewError(utility.ERROR_REQUEST_CODE, utility.ERROR_REQUEST_MSG)

query:
	db := controller.GetDB()
	//查询post
	db.Where(&article).First(&article)
	//查询tags
	db.Model(&article).Related(&article.Tags, "tags")
	//db.Model(&article).Association("tags").Find(&article.Tags)

	if db.Error != nil {
		return nil, db.Error
	}
	res := model.RESGetArticle{
		Aid:    article.ID,
		Title:  article.Title,
		Author: article.AuthorName,
		ReplayCount:article.ReplayCount,
		Tags: func(tags []*model.Tag) []string {
			ts := []string{}
			for i, _ := range tags {
				ts = append(ts, tags[i].Name)
			}
			return ts
		}(article.Tags),
		Context:        article.Context,
		CreateDatetime: formatDatetime(article.CreatedAt),
		EditDatetime:   formatDatetime(article.UpdatedAt),
	}
	return &res, nil
}
func DelArticle(req *model.REQDelArticle) (err error) {

	article := model.Article{}
	article.Title = req.Title
	tx := controller.GetDB().Begin()

	if err = tx.Model(&article).Where(&article).Delete(&article).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func PostReplay(req *model.REQNewReplay) (err error) {
	if isNullOrEmpty(req.Title) || isNullOrEmpty(req.Context) {
		return utility.NewError(utility.ERROR_REQUEST_CODE, utility.ERROR_REQUEST_MSG)
	}
	article := model.Article{}

	if !isNullOrEmpty(req.Aid) {
		aid, err := parseID(req.Aid)
		if err != nil {
			return err
		}
		article.ID = aid
	}
	user := req.CurrUser
	article.Title = req.Title

	db := controller.GetDB()
	if err = db.Model(&article).Where(&article).Select("id,title,replay_count").First(&article).Error; err != nil {
		return err
	}
	replay := model.Replay{
		ArticleTitle: article.Title,
		AuthorName:   user.Name,
		Context:      req.Context,
		Count:article.ReplayCount+1,
	}
	tx := db.Begin()
	if err=tx.Model(&article).Where(&article).Update(article).Error;err!=nil{
		tx.Rollback()
		return err
	}
	if err = tx.Save(&replay).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func GetArticleReplays(req *model.REQGetReplays) (*model.RESGetReplays, error) {

	article := &model.Article{
		Title: req.Title,
	}
	db := controller.GetDB()
	if err := db.Model(&article).Where(&article).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&article).
		Select("id,author_name,context,created_at").
		Order("count desc").
		Related(&article.Replays, "article_title").Error;
		err != nil {
		return nil, err
	}
	res := model.RESGetReplays{}
	res.Aid = article.ID
	res.ArticleTitle = article.Title

	for i, _ := range article.Replays {
		v := &article.Replays[i]
		res.Replays = append(res.Replays, struct {
			Rid            uint   `json:"rid"`
			Username       string `json:"username"`
			Context        string `json:"context"`
			CreateDatetime string `json:"create_datetime"`
		}{Rid: v.ID, Username: v.AuthorName, Context: v.Context, CreateDatetime: formatDatetime(v.CreatedAt)})
	}
	return &res, nil
}
func DelArticleReplay(req *model.REQDelReplays)(err error){
	rid,err:=parseID(req.Rid)
	if err!=nil{
		return err
	}
	replay:=model.Replay{}
	replay.ArticleTitle=req.Title
	replay.ID=rid
	//todo
	return nil
}
func GetUserInfo(req *model.REQGetUserInfo) (*model.RESGetUserInfo, error) {
	var err error
	if isNullOrEmpty(req.Uid) && isNullOrEmpty(req.Username) {
		return nil, utility.NewError(utility.ERROR_REQUEST_CODE, utility.ERROR_REQUEST_MSG)
	}
	quser := model.User{}

	if !isNullOrEmpty(req.Uid) {
		uid, err := parseID(req.Uid)
		if err != nil {
			return nil, err
		}
		quser.ID = uid
	}
	quser.Name = req.Username

	db := controller.GetDB()

	if err = db.Model(&quser).Where(&quser).First(&quser).Error; err != nil {
		return nil, err
	}

	if err = db.Model(&quser).Select("title,created_at").Order("created_at asc").Related(&quser.Articles, "author_id").Error; err != nil {
		return nil, err
	}
	if err = db.Model(&model.User{}).Select("article_title,context,created_at").Order("created_at asc").Related(&quser.Replays, "author_name").Error; err != nil {
		return nil, err
	}

	res := model.RESGetUserInfo{}

	for i, _ := range quser.Articles {
		v := &quser.Articles[i]

		res.PostArticle = append(res.PostArticle, struct {
			Title          string `json:"title"`
			CreateDatetime string `json:"create_datetime"`
		}{Title: v.Title, CreateDatetime: formatDatetime(v.CreatedAt)})
	}

	for i, _ := range quser.Replays {
		v := &quser.Replays[i]

		res.PostReplay = append(res.PostReplay, struct {
			Title          string `json:"title"`
			Replay         string `json:"replay"`
			CreateDatetime string `json:"create_datetime"`
		}{Title: v.ArticleTitle, Replay: v.Context, CreateDatetime: formatDatetime(v.CreatedAt)})
	}

	res.Username = quser.Name
	res.Uid = quser.ID
	res.Email = quser.Email
	res.Permission = quser.Permission
	res.CreateDatetime = formatDatetime(quser.CreatedAt)
	res.UpdateDatetime = formatDatetime(quser.UpdatedAt)
	return &res, nil
}
