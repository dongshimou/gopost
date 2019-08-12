package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopost/src/base"
	"gopost/src/logger"
	"gopost/src/model"
	"log"
	"time"
)

var (
	db *gorm.DB
)

func Get() *gorm.DB {
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

	db = db.AutoMigrate(
		&model.Article{},
		&model.Replay{},
		&model.User{},
		&model.Tag{},
		&model.Stat{},
		&model.Mood{},
	)
	if logger.DEBUG {
		db.LogMode(true)
	}
	return db.Error
}
