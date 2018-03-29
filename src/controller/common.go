package controller

import (
	_ "github.com/go-sql-driver/mysql"
	"base"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"model"
)

var (
	db *gorm.DB
)

func InitDB() error {
	var err error
	cfg := base.GetConfig().Database
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err = gorm.Open("mysql", args)
	if err != nil {
		log.Print(err)
		return err
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	if db.HasTable(&model.Post{}) == false {
		db.CreateTable(&model.Post{})
	}
	if db.HasTable(&model.Replay{}) == false {
		db.CreateTable(&model.Replay{})
	}
	if db.HasTable(&model.User{}) == false {
		db.CreateTable(&model.User{})
	}
	if db.HasTable(&model.Tag{}) == false {
		db.CreateTable(&model.Tag{})
	}
	return nil
}
