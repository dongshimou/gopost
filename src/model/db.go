package model

import (
	"github.com/jinzhu/gorm"
)

const (
	DB_id         = "id"
	DB_created_at = "created_at"
	DB_updated_at = "updated_at"
	DB_deleted_at = "deleted_at"
	DB_asc		="asc"
	DB_desc		="desc"
)

type Article struct {
	gorm.Model
	Title             string   `gorm:"size:255;not null;unique;"`
	Author            User     `gorm:"FOREIGNKEY:AuthorName;"`
	AuthorName        string   `gorm:"size:255;not null;index;"`
	ReplayCount       uint     `gorm:"default:0"`
	Tags              []*Tag   `gorm:"many2many:tag_articles;"`
	Context           string   `gorm:"size:65535;"`
	Replays           []Replay `gorm:"FOREIGNKEY:ArticleTitle;"`
	PermissionRequire int      `gorm:"default:1;"`
}

const (
	Table_Article_Title             = "title"
	Table_Article_AuthorName        = "author_name"
	Table_Article_ReplayCount       = "replay_count"
	Table_Article_Context           = "context"
	Table_Article_PermissionRequire = "permission_require"
)

type Replay struct {
	gorm.Model
	Article      Article `gorm:"FOREIGNKEY:ArticleTitle;"`
	ArticleTitle string  `gorm:"size:255;not null;index;"` //id

	Author     User   `gorm:"FOREIGNKEY:AuthorName;"`
	AuthorName string `gorm:"size:255;not null;index;"` //id
	Context    string `gorm:"size:2048;"`
	Count      uint   `gorm:"not null;index;"`
}
const(
	Table_Replay_ArticleTitle ="article_title"
	Table_Replay_AuthorName   ="author_name"
	Table_Replay_Context      ="context"
	Table_Replay_Count        ="count"
)
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
	Articles   []Article `gorm:"FOREIGNKEY:AuthorID;"`
	Replays    []Replay  `gorm:"FOREIGNKEY:AuthorName;"`

	SignInIP string `gorm:"size:128"`
	SignUpIP string `gorm:"size:128"`
}

const (
	Replay_Read   = 1
	Replay_Create = 1 << 1
	Replay_Delete = 1 << 2
	Replay_Update = 1 << 3

	Article_Read   = 1 << 4
	Article_Create = 1 << 5
	Article_Delete = 1 << 6
	Article_Update = 1 << 7

	User_Read   = 1 << 8
	User_Create = 1 << 9
	User_Delete = 1 << 10
	User_Update = 1 << 11
)
