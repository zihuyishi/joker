package model

type User struct {
	tableName struct{} `sql:"user"`
	Id int64 `json:"id" sql:"id,pk" form:"id"`
	Name string `json:"name" sql:"name" form:"name"`
	Password string	`json:"password" sql:"password" form:"password"`
}
