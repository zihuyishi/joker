package dao

import "github.com/zihuyishi/joker/web/model"

func (d *Dao) InsertJoker(joker *model.Joker) error {
	err := d.db.Insert(joker)
	return err
}

func (d *Dao) FindById(id int64) (*model.Joker, error) {
	joker := &model.Joker{
		Id: id,
	}
	err := d.db.Select(joker)
	return joker, err
}