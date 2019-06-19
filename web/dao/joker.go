package dao

import (
	"fmt"
	"github.com/zihuyishi/joker/web/model"
)

func (dao *Dao) InsertJoker(joker *model.Joker) error {
	err := dao.db.Insert(joker)
	return err
}

func (dao *Dao) FindJokerById(id int64) (*model.Joker, error) {
	joker := &model.Joker{
		Id: id,
	}
	err := dao.db.Select(joker)
	return joker, err
}

func (dao *Dao) RandomJoker(count int) (*[]model.Joker, error) {
	var jokers []model.Joker
	sql := fmt.Sprintf("SELECT * FROM joker ORDER BY random() LIMIT %d", count)
	_, err := dao.db.Model((*model.Joker)(nil)).Query(&jokers, sql)
	return &jokers, err
}