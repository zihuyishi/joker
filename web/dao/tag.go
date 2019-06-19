package dao

import (
	"github.com/go-pg/pg"
	"github.com/zihuyishi/joker/web/model"
)

func (dao *Dao) InsertTag(tag *model.Tag) error {
	err := dao.db.Insert(tag)
	return err
}

func (dao *Dao) FindTagById(id int64) (*model.Tag, error) {
	tag := &model.Tag{
		Id: id,
	}
	err := dao.db.Select(tag)
	return tag, err
}

func (dao *Dao) FindTagsByIds(ids []int64) (*[]model.Tag, error) {
	var tags *[]model.Tag
	err := dao.db.Model(tags).Where("id IN (?)", pg.In(ids)).Select()
	return tags, err
}

func (dao *Dao) FindJokerTags(jokerId int64) (*[]model.Tag, error) {
	var jokerTags []model.JokerTag
	err := dao.db.Model(&jokerTags).Where("joker_id = ?", jokerId).Select()
	if err != nil {
		return nil, err
	}
	tagIds := make([]int64, len(jokerTags))
	for i := 0; i < len(jokerTags); i++  {
		tagIds[i] = jokerTags[i].Id
	}
	var tags []model.Tag
	err = dao.db.Model(&tags).Where("id IN (?)", pg.In(tagIds)).Select()
	return &tags, err
}