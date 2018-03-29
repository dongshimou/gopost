package model

import "time"

type Post struct {
	Pid               int64  `gorm:"primary_key;auto_increment;"`
	Title             string `gorm:"size:255;not null;"`
	Author            User
	Tags              []Tag
	Replays           []Replay
	Context           string    `gorm:"size:65535"`
	CreateDatetime    time.Time `gorm:"default:now();"`
	EditDatetime      time.Time `gorm:"default:now();"`
	EditPersion       User
	PermissionRequire int
}
type Replay struct {
	Rid            int64 `gorm:"primary_key;auto_increment;"`
	Post           Post
	User           User
	Context        string    `gorm:"size:2048"`
	CreateDatetime time.Time `gorm:"default:now();"`
}
type Tag struct {
	Tid            int64     `gorm:"primary_key;auto_increment;"`
	Name           string    `gorm:"size:255"`
	CreateDatetime time.Time `gorm:"default:now();"`
}

type User struct {
	Uid            int64     `gorm:"primary_key;auto_increment;"`
	Name           string    `gorm:"size:255;not null;"`
	Email          string    `gorm:"size:255"`
	Permission     int       `gorm:"default:1"`
	CreateDatetime time.Time `gorm:"default:now();"`
	Token          string    `gorm:"size:2048"`
	Posts          []Post
	Replays        []Replay
}
