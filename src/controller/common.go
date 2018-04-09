package controller

import (
	"base"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"logger"
	"model"
	"time"
)

var (
	db *gorm.DB
)

func GetDB() *gorm.DB {
	return db
}
func InitDB() error {
	var err error
	cfg := base.GetConfig().Database
	logger.Print(cfg)
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err = gorm.Open("mysql", args)
	if err != nil {
		log.Print(err)
		return err
	}

	//连接池配置
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	//使用默认表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")

	if db.HasTable(&model.Article{}) == false {
		db = db.CreateTable(&model.Article{})
	}
	if db.HasTable(&model.Replay{}) == false {
		db = db.CreateTable(&model.Replay{})
	}
	if db.HasTable(&model.User{}) == false {
		db = db.CreateTable(&model.User{})
	}
	if db.HasTable(&model.Tag{}) == false {
		db = db.CreateTable(&model.Tag{})
	}
	return db.Error
}
