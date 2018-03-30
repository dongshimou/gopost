package model

import "time"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type RESLogin struct {
}

type RESGetPost struct {
	Title          string    `json:"title"`
	Author         string    `json:"author"`
	Tags           []string  `json:"tags"`
	Context        string    `json:"context"`
	CreateDatetime time.Time `json:"create_datetime" time_format:"2006-01-02 15:04:05" time_utc:"1" `
	EditDatetime   time.Time `json:"edit_datetime" time_format:"2006-01-02 15:04:05" time_utc:"1" `
}
