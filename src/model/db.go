package model

import (
	"github.com/jinzhu/gorm"
)

const (
	DB_created_at = "created_at"
	DB_updated_at = "updated_at"
	DB_deleted_at = "deleted_at"
)

type Article struct {
	gorm.Model
	Title    string `gorm:"size:255;not null;unique;"`
	Author   User   `gorm:"FOREIGNKEY:AuthorID;"`
	AuthorID uint

	Editor   User `gorm:"FOREIGNKEY:EditorID;"`
	EditorID uint

	Tags              []*Tag   `gorm:"many2many:tag_articles;"`
	Context           string   `gorm:"size:65535;"`
	Replays           []Replay `gorm:"FOREIGNKEY:ArticleTitle;"`
	PermissionRequire int      `gorm:"default:1;"`
}

type Replay struct {
	gorm.Model
	Article      Article `gorm:"FOREIGNKEY:ArticleTitle;"`
	ArticleTitle string  `gorm:"size:255;not null;index;"` //id

	Author     User   `gorm:"FOREIGNKEY:AuthorName;"`
	AuthorName string `gorm:"size:255;not null;index;"` //id
	Context    string `gorm:"size:2048;"`
}
type Tag struct {
	gorm.Model
	Name     string     `gorm:"size:255;unique;"`
	Articles []*Article `gorm:"many2many:tag_articles;"`
}

type User struct {
	gorm.Model
	Name       string    `gorm:"size:255;not null;unique;"`
	Email      string    `gorm:"size:255;unique;"`
	Password   string    `gorm:"size:128;"`
	Permission int       `gorm:"default:1;"`
	Token      string    `gorm:"size:2048;"`
	Articles   []Article `gorm:"FOREIGNKEY:AuthorID;"`
	Replays    []Replay  `gorm:"FOREIGNKEY:AuthorName;"`
}

const (
	Article_PermissionRead   = 0x0000
	Article_PermissionCreate = 0x0010
	Replay_PermissionRead    = 0x0000
	Replay_PermissionCreate  = 0x0001
	User_PermissionRead = 0x0000
	User_PermissionCreate = 0x0100
)
