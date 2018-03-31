package model

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title    string `gorm:"size:255;not null;unique;"`
	Author   User   `gorm:"FOREIGNKEY:AuthorID;"`
	AuthorID uint

	Editor   User `gorm:"FOREIGNKEY:EditorID;"`
	EditorID uint

	Tags              []*Tag   `gorm:"many2many:tag_posts;"`
	Context           string   `gorm:"size:65535;"`
	Replays           []Replay `gorm:"FOREIGNKEY:PostID;"`
	PermissionRequire int      `gorm:"default:1;"`
}

const (
	DB_Post_Tags = "tags"
)

type Replay struct {
	gorm.Model
	PostID  uint
	UserID  uint
	Context string `gorm:"size:2048;"`
}
type Tag struct {
	gorm.Model
	Name  string  `gorm:"size:255;unique;"`
	Posts []*Post `gorm:"many2many:tag_posts;"`
}

const (
	DB_Table_Tag = "tags"
)

type User struct {
	gorm.Model
	Name       string   `gorm:"size:255;not null;unique;"`
	Email      string   `gorm:"size:255;unique;"`
	Password   string   `gorm:"size:128;"`
	Permission int      `gorm:"default:1;"`
	Token      string   `gorm:"size:2048;"`
	Posts      []Post   `gorm:"FOREIGNKEY:AuthorID;"`
	Replays    []Replay `gorm:"FOREIGNKEY:UserID;"`
}
