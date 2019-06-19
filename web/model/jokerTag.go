package model

type JokerTag struct {
	tableName struct{} `sql:"joker_tag"`
	Id int64 `sql:"id,pk"`
	TagId int64	`sql:"tag_id"`
	JokerId int64 `sql:"joker_id"`
}
