package dao

import "github.com/zihuyishi/joker/web/model"

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