package model

import "time"

type Joker struct {
	tableName struct{} `sql:"joker"`
	Id int64 `json:"id" sql:"id,pk" form:"id"`
	Title string `json:"title" sql:"title" form:"title"`
	Content string `json:"content" sql:"content" form:"content"`
	Time time.Time `json:"time" sql:"time" form:"-"`
	Tags *[]Tag `json:"tag" sql:"-" form:"-"`
}
