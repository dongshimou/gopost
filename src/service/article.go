package service

import (
	"controller"
	"model"
	"utility"
)

func PostNewArticle(req *model.REQNewArticle) error {

	user := model.User{
		Name: "root",
	}
	db := controller.GetDB()

	tags := []*model.Tag{}
	for i, _ := range req.Tags {
		t := model.Tag{Name: req.Tags[i]}
		db.FirstOrCreate(&t, &t)
		tags = append(tags, &t)
	}

	if db.Where(&user).First(&user).Error != nil {
		return db.Error
	}
	//验证权限
	if err:=utility.VerifyPermission(user.Permission,model.Article_PermissionCreate);err!=nil{
	return err
	}

	post := model.Article{
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
	//查询user
	db.Model(&article).Related(&article.Author, "author_id")

	if db.Error != nil {
		return nil, db.Error
	}
	res := model.RESGetArticle{
		Aid:    article.ID,
		Title:  article.Title,
		Author: article.Author.Name,
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

func NewReplay(req *model.REQNewReplay) (err error) {
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
	user := model.User{
		Name: "root",
	}
	article.Title = req.Title

	db := controller.GetDB()
	if err = db.Model(&article).Where(&article).Select("title").First(&article).Error; err != nil {
		return err
	}
	if err = db.Model(&user).Where(&user).First(&user).Error; err != nil {
		return err
	}
	replay := model.Replay{
		ArticleTitle: article.Title,
		AuthorName:   user.Name,
		Context:      req.Context,
	}
	tx := db.Begin()
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
	if err := db.Model(&article).Related(&article.Replays, "article_title").Error; err != nil {
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

func GetUserInfo(req *model.REQGetUserInfo) (*model.RESGetUserInfo, error) {
	var err error
	if isNullOrEmpty(req.Uid) && isNullOrEmpty(req.Username) {
		return nil, utility.NewError(utility.ERROR_REQUEST_CODE, utility.ERROR_REQUEST_MSG)
	}
	user := model.User{}

	if !isNullOrEmpty(req.Uid) {
		uid, err := parseID(req.Uid)
		if err != nil {
			return nil, err
		}
		user.ID = uid
	}
	user.Name = req.Username

	db := controller.GetDB()

	if err = db.Model(&user).Where(&user).First(&user).Error; err != nil {
		return nil, err
	}

	if err = db.Model(&user).Select("title,created_at").Related(&user.Articles, "author_id").Error; err != nil {
		return nil, err
	}
	if err = db.Model(&model.User{}).Related(&user.Replays, "author_name").Error; err != nil {
		return nil, err
	}

	res := model.RESGetUserInfo{}

	for i, _ := range user.Articles {
		v := &user.Articles[i]

		res.PostArticle = append(res.PostArticle, struct {
			Title          string `json:"title"`
			CreateDatetime string `json:"create_datetime"`
		}{Title: v.Title, CreateDatetime: formatDatetime(v.CreatedAt)})
	}

	for i, _ := range user.Replays {
		v := &user.Replays[i]

		res.PostReplay = append(res.PostReplay, struct {
			Title          string `json:"title"`
			Replay         string `json:"replay"`
			CreateDatetime string `json:"create_datetime"`
		}{Title: v.ArticleTitle, Replay: v.Context, CreateDatetime: formatDatetime(v.CreatedAt)})
	}

	res.Username = user.Name
	res.Uid = user.ID
	res.Email = user.Email
	res.Permission = user.Permission
	res.CreateDatetime = formatDatetime(user.CreatedAt)
	res.UpdateDatetime = formatDatetime(user.UpdatedAt)
	return &res, nil
}
