package model

import "time"

type Joker struct {
	tableName struct{} `sql:"joker"`
	Id int64 `json:"id" sql:"id,pk"`
	Title string `json:"title" sql:"title"`
	Content string `json:"content" sql:"content"`
	Time time.Time `json:"time" sql:"time"`
}
