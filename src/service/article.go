package service

import (
	"controller"
	"model"
	"utility"
	"logger"
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
		Title:      req.Title,
		Context:    req.Context,
		Tags:       tags,
		AuthorName: user.Name,
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
	if err := db.Where(&article).First(&article).Error; err != nil {
		return nil, err
	}
	//查询tags
	if err := db.Model(&article).Related(&article.Tags, "tags").Error; err != nil {
		//db.Model(&article).Association("tags").Find(&article.Tags)
		return nil, err
	}
	next := model.Article{}
	prev := model.Article{}
	//上一篇和下一篇
	db.Model(&prev).Where("id < ?",article.ID).Select("title").Last(&prev)
	db.Model(&next).Where("id > ?",article.ID).Select("title").First(&next)
	res := model.RESGetArticle{
		Aid:         article.ID,
		Title:       article.Title,
		Author:      article.AuthorName,
		ReplayCount: article.ReplayCount,
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
		Next:           next.Title,
		Prev:           prev.Title,
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
	logger.Print(buildArgs(",", model.DB_id, model.Table_Article_Title, model.Table_Article_ReplayCount))
	if err = db.Model(&article).
		Where(&article).
		Select(buildArgs(",", model.DB_id, model.Table_Article_Title, model.Table_Article_ReplayCount)).
		First(&article).Error; err != nil {
		return err
	}
	replay := model.Replay{
		ArticleTitle: article.Title,
		AuthorName:   user.Name,
		Context:      req.Context,
		Count:        article.ReplayCount + 1,
	}
	tx := db.Begin()
	if err = tx.Save(&replay).Error; err != nil {
		tx.Rollback()
		return err
	}
	article.ReplayCount+=1
	where:=model.Article{}
	where.ID=article.ID
	if err = tx.Model(&article).Where(&where).Update(article).Error; err != nil {
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
		Select(buildArgs(",", model.DB_id, model.Table_Article_AuthorName, model.Table_Article_Context, model.DB_created_at)).
		Order(buildArgs(" ", model.Table_Replay_Count, model.DB_desc)).
		Related(&article.Replays, "article_title").Error; err != nil {
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
func DelArticleReplay(req *model.REQDelReplays) (err error) {
	replay := model.Replay{}

	if !isNullOrEmpty(req.Rid) {
		rid, err := parseID(req.Rid)
		if err != nil {
			return err
		}
		replay.ID = rid
	} else {
		count, err := parseCount(req.Count)
		if err != nil {
			return err
		}
		replay.Count = count
		replay.ArticleTitle = req.Title
	}
	tx := controller.GetDB().Begin()
	if err = tx.Model(&replay).Where(&replay).Delete(&replay).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
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

	if err = db.Model(&quser).
		Select(buildArgs(",", model.Table_Article_Title, model.DB_created_at)).
		Order(buildArgs(" ", model.DB_created_at, model.DB_asc)).
		Related(&quser.Articles, "author_id").Error; err != nil {
		return nil, err
	}
	if err = db.Model(&model.User{}).
		Select(buildArgs(",", model.Table_Replay_ArticleTitle, model.Table_Replay_Context, model.DB_created_at)).
		Order(buildArgs(" ", model.DB_created_at, model.DB_asc)).
		Related(&quser.Replays, "author_name").Error; err != nil {
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
