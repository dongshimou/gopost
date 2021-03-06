package service

import (
	"database/sql"
	"errors"
	"github.com/jinzhu/gorm"
	"gopost/src/logger"
	"gopost/src/model"
	"gopost/src/orm"
	"gopost/src/protocol"
	"gopost/src/utility"
	"strings"
	"time"
)

func CreateArticle(req *protocol.REQNewArticle) error {

	db := orm.Get().Begin()
	if err := createOrupdateArticle(db, req, ""); err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	share2sns(req.SNS)
	return nil
}

func share2sns(snsList string) error {
	slist := strings.Split(snsList, ",")
	for _, sns := range slist {
		switch sns {
		case "twitter":

		default:
			logger.Debug("not support method!")
		}
	}
	return nil
}
func createOrupdateArticle(tx *gorm.DB, req *protocol.REQNewArticle, oldTitle string) error {
	tags := []*model.Tag{}
	for i, _ := range req.Tags {
		t := model.Tag{Name: req.Tags[i]}
		if err := tx.FirstOrCreate(&t, &t).Error; err != nil {
			return err
		}
		tags = append(tags, &t)
	}
	post := model.Article{
		Title:      req.Title,
		Context:    req.Context,
		Tags:       tags,
		AuthorName: req.CurrUser.Name,
	}
	if oldTitle == "" {
		logger.Debug("create new ", req.Title)
		if err := tx.Save(&post).Error; err != nil {
			return err
		}
		logger.Debug("create", req.Title, "success")
	} else {
		art := model.Article{}
		logger.Debug("update -> ", oldTitle)
		if err := tx.Model(&model.Article{}).Where(&model.Article{Title: oldTitle}).Last(&art).Error; err != nil {
			return err
		}
		if err := tx.Set("gorm:save_associations", true).Model(&model.Article{}).Where(&model.Article{Title: oldTitle}).Update(&post).Error; err != nil {
			return err
		}
		//gorm bug : can't del tags
		{
			tagsRes := []model.TagArticles{}
			if err := tx.Model(&model.TagArticles{}).Where(&model.TagArticles{ArticleId: art.ID}).Find(&tagsRes).Error; err != nil {
				if err == sql.ErrNoRows {

				}
				return err
			}
			//删除
			for _, v := range tagsRes {
				find := false
				for _, vv := range post.Tags {
					if v.TagId == vv.ID {
						find = true
						break
					}
				}
				if !find {
					tx.Where(&v).Delete(&v)
					//todo 如果tag已经没有使用了,将其从tags里删除
				}
			}
			//新增
			for _, v := range post.Tags {
				find := false
				for _, vv := range tagsRes {
					if v.ID == vv.TagId {
						find = true
						break
					}
				}
				if !find {
					tx.Create(&model.TagArticles{ArticleId: art.ID, TagId: v.ID})
				}
			}
		}

		logger.Debug("update", req.Title, "success")
	}
	return nil
}
func UpdateArticle(req *protocol.REQUpdateArticle) error {
	db := orm.Get().Begin()
	if err := createOrupdateArticle(db, &req.REQNewArticle, req.OldTitle); err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}
func GetStat(req *protocol.REQGetStat) (*protocol.RESGetStat, error) {
	res := protocol.RESGetStat{}

	//倒序统计ip
	sql := orm.Get().Model(&model.Stat{}).
		Select("date,count(ip) as count").
		Group("date").Order("date desc")

	if req.Date != "" {
		sql = sql.Having("date=?", req.Date)
	}
	if err := sql.Scan(&res.List).
		Error; err != nil {
		return nil, err
	}
	return &res, nil
}
func StatIp(ip string) error {
	db := orm.Get()
	date := utility.FormatDate(time.Now())
	stat := model.Stat{}
	stat.Date = date
	stat.Ip = ip
	count := 0
	if err := db.Model(&model.Stat{}).Where(&stat).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return db.Create(&stat).Error
	} else {
		return nil
	}
}
func GetArticles(req *protocol.REQGetArticles) (*protocol.RESGetArticles, error) {

	t1, errT1 := parseTime(req.Time)
	t2, errT2 := parseUnix(req.Time)
	if errT1 != nil && errT2 != nil {
		return nil, utility.NewError(utility.ERROR_REQUEST_CODE, utility.ERROR_REQUEST_MSG)
	}
	var befor time.Time
	if errT1 == nil {
		befor = t1
	} else {
		befor = t2
	}
	logger.Debug(befor.Unix())
	logger.Debug(befor.String())
	logger.Debug(formatDatetime(befor))

	limit, err := parseCount(req.Size)
	if err != nil {
		return nil, err
	}
	arts := []model.Article{}

	db := orm.Get()
	if err = db.Model(&model.Article{}).
		Where("created_at < ?", befor).
		Order("created_at desc").
		Limit(limit).
		Find(&arts).Error; err != nil {
		return nil, err
	}

	for i, _ := range arts {
		if err = db.Model(&arts[i]).Select("name").
			Related(&arts[i].Tags, "tags").Error; err != nil {
			return nil, err
		}
	}

	resData := make([]protocol.RESGetArticle, len(arts))
	for i := 0; i < len(arts); i++ {
		a := &arts[i]
		resData[i] = protocol.RESGetArticle{
			a.ID,
			a.Title,
			a.AuthorName,
			func(tags []*model.Tag) []string {
				ts := []string{}
				for i, _ := range tags {
					ts = append(ts, tags[i].Name)
				}
				return ts
			}(a.Tags),
			a.Context,
			formatDatetime(a.CreatedAt),
			formatDatetime(a.UpdatedAt),
			"",
			"",
		}
	}
	return &protocol.RESGetArticles{resData}, nil
}
func GetArticle(req *protocol.REQGetArticle) (*protocol.RESGetArticle, error) {
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
	db := orm.Get()
	//查询post
	if err := db.Where(&article).First(&article).Error; err != nil {
		return nil, err
	}
	//查询tags
	if err := db.Model(&article).Select("name").
		Related(&article.Tags, "tags").Error; err != nil {
		//db.Model(&article).Association("tags").Find(&article.Tags)
		return nil, err
	}
	next := model.Article{}
	prev := model.Article{}
	//上一篇和下一篇
	db.Model(&prev).Where("id < ?", article.ID).Select("title").Last(&prev)
	//db.Model(&prev).Where("created_at < ?", article.CreatedAt).Select("title").Last(&prev)
	db.Model(&next).Where("id > ?", article.ID).Select("title").First(&next)
	//db.Model(&next).Where("created_at > ?", article.CreatedAt).Select("title").First(&next)
	res := protocol.RESGetArticle{
		Aid:    article.ID,
		Title:  article.Title,
		Author: article.AuthorName,
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
func DelArticle(req *protocol.REQDelArticle) (err error) {

	article := model.Article{}
	article.Title = req.Title
	tx := orm.Get().Begin()
	if err = tx.Model(&article).Where(&article).Delete(&article).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func GetTags(req *protocol.REQGetTags) (*protocol.RESGetTags, error) {

	db := orm.Get()
	art := model.Article{Title: req.Title}
	if err := db.Model(&model.Article{}).Where(&art).Last(&art).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&art).
		Select("name").Related(&art.Tags, "tags").Error; err != nil {
		return nil, err
	}
	res := protocol.RESGetTags{}
	for _, v := range art.Tags {
		res.Names = append(res.Names, v.Name)
	}
	return &res, nil
}
func GetAllTags() (*protocol.RESGetTags, error) {
	db := orm.Get()

	tags := []model.Tag{}
	if err := db.Model(&model.Tag{}).Select("name").Find(&tags).Error; err != nil {
		return nil, err
	}
	res := protocol.RESGetTags{}
	for _, v := range tags {
		res.Names = append(res.Names, v.Name)
	}
	return &res, nil
}
func PostReplay(req *protocol.REQNewReplay) (err error) {
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

	db := orm.Get()
	logger.Print(buildArgs(",", model.DB_id, model.Table_Article_Title))
	if err = db.Model(&article).
		Where(&article).
		Select(buildArgs(",", model.DB_id, model.Table_Article_Title)).
		First(&article).Error; err != nil {
		return err
	}
	replay := model.Replay{
		ArticleTitle: article.Title,
		AuthorName:   user.Name,
		Context:      req.Context,
		IpAddress:    req.IpAddress,
	}
	tx := db.Begin()
	if err = tx.Save(&replay).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func GetArticleReplays(req *protocol.REQGetReplays) (*protocol.RESGetReplays, error) {

	article := &model.Article{
		Title: req.Title,
	}
	db := orm.Get()
	if err := db.Model(&article).Select(model.Table_Article_Title).
		Where(&article).First(&article).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&article).
		Select(strings.Join([]string{
			model.DB_id,
			model.Table_Replay_AuthorName,
			model.Table_Replay_IpAddress,
			model.Table_Replay_Context,
			model.DB_created_at,
		}, ",")).
		Order(strings.Join([]string{model.DB_created_at, model.DB_desc}, " ")).
		Related(&article.Replays, "Replays").Error; err != nil {
		return nil, err
	}
	res := protocol.RESGetReplays{}
	res.Aid = article.ID
	res.ArticleTitle = article.Title

	for i, _ := range article.Replays {
		v := &article.Replays[i]
		res.Replays = append(res.Replays, protocol.RESGetReplaysSingle{
			Rid:            v.ID,
			Username:       v.AuthorName,
			Context:        v.Context,
			IpAddress:      v.IpAddress,
			CreateDatetime: formatDatetime(v.CreatedAt),
		})
	}
	return &res, nil
}
func DelArticleReplay(req *protocol.REQDelReplays) (err error) {
	replay := model.Replay{}

	if !isNullOrEmpty(req.Rid) {
		rid, err := parseID(req.Rid)
		if err != nil {
			return err
		}
		replay.ID = rid
	} else {
		return errors.New("it's not a rid")
	}
	tx := orm.Get().Begin()
	if err = tx.Model(&model.Replay{}).Where(&replay).Delete(&replay).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func GetUserInfo(req *protocol.REQGetUserInfo) (*protocol.RESGetUserInfo, error) {
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

	db := orm.Get()

	if err = db.Model(&quser).Where(&quser).First(&quser).Error; err != nil {
		return nil, err
	}

	if err = db.Model(&quser).
		Select(buildArgs(",", model.Table_Article_Title, model.DB_created_at)).
		Order(buildArgs(" ", model.DB_created_at, model.DB_asc)).
		Related(&quser.Articles, "Articles").Error; err != nil {
		return nil, err
	}
	if err = db.Model(&model.User{}).
		Select(buildArgs(",", model.Table_Replay_ArticleTitle,
			model.Table_Replay_Context, model.DB_created_at)).
		Order(buildArgs(" ", model.DB_created_at, model.DB_asc)).
		Related(&quser.Replays, "Replays").Error; err != nil {
		return nil, err
	}

	res := protocol.RESGetUserInfo{}

	for i, _ := range quser.Articles {
		v := &quser.Articles[i]

		res.PostArticle = append(res.PostArticle, protocol.RESGetUserInfoArticle{
			Title:          v.Title,
			CreateDatetime: formatDatetime(v.CreatedAt),
		})
	}

	for i, _ := range quser.Replays {
		v := &quser.Replays[i]

		res.PostReplay = append(res.PostReplay, protocol.RESGetUserInfoReplay{
			Title:          v.ArticleTitle,
			Replay:         v.Context,
			CreateDatetime: formatDatetime(v.CreatedAt),
		})
	}

	res.Username = quser.Name
	res.Uid = quser.ID
	res.Email = quser.Email
	res.Permission = quser.Permission
	res.CreateDatetime = formatDatetime(quser.CreatedAt)
	res.UpdateDatetime = formatDatetime(quser.UpdatedAt)
	return &res, nil
}
