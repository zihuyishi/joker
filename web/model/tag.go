package model

type Tag struct {
	tableName struct{} `sql:"tag"`
	Id int64 `json:"id" sql:"id,pk"`
	Name string `json:"name" sql:"name"`
}