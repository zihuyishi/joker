package dao

import (
	"github.com/go-pg/pg"
)

type Dao struct {
	db *pg.DB
}

func New(config *pg.Options) *Dao {
	db := pg.Connect(config)
	dao := Dao {
		db,
	}
	return &dao
}
