package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title    string `gorm:"size:255;not null;unique;"`
	Author   User   `gorm:"FOREIGNKEY:AuthorID;"`
	AuthorID uint

	Editor   User `gorm:"FOREIGNKEY:EditorID;"`
	EditorID uint

	Tags              []*Tag   `gorm:"many2many:tag_posts;"`
	Context           string   `gorm:"size:65535;"`
	Replays           []Replay `gorm:"FOREIGNKEY:ArticleID;"`
	PermissionRequire int      `gorm:"default:1;"`
}

type Replay struct {
	gorm.Model
	ArticleID uint
	UserID    uint
	Context   string `gorm:"size:2048;"`
}
type Tag struct {
	gorm.Model
	Name    string     `gorm:"size:255;unique;"`
	Article []*Article `gorm:"many2many:tag_posts;"`
}

type User struct {
	gorm.Model
	Name       string    `gorm:"size:255;not null;unique;"`
	Email      string    `gorm:"size:255;unique;"`
	Password   string    `gorm:"size:128;"`
	Permission int       `gorm:"default:1;"`
	Token      string    `gorm:"size:2048;"`
	Article    []Article `gorm:"FOREIGNKEY:AuthorID;"`
	Replays    []Replay  `gorm:"FOREIGNKEY:UserID;"`
}
