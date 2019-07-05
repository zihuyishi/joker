package dao

import (
	"crypto/md5"
	"fmt"
	"github.com/zihuyishi/joker/web/code"
	"github.com/zihuyishi/joker/web/model"
)

func (dao *Dao) FindUserById(id int64) (*model.User, error) {
	user := &model.User{
		Id: id,
	}
	err := dao.db.Select(user)
	return user, err
}

func (dao *Dao) InsertUser(name string, password string) (*model.User, error) {
	hashPass := hashPassword(password)
	user := &model.User{
		Name: name,
		Password: hashPass,
	}
	err := dao.db.Insert(user)
	return user, err
}


func (dao *Dao) VerifyUserPassword(name string, password string) (*model.User, error) {
	user := &model.User{
		Name: name,
	}
	err := dao.db.Model(user).Where("name = ?", name).Select()
	if err != nil {
		return nil, err
	}
	hashPass := hashPassword(password)
	result := hashPass == user.Password
	if !result {
		return nil, code.ERROR_WRONG_PASSWORD
	}
	return user, nil
}

func (dao *Dao) ChangePassword(uid int64, newPass string) error {
	hashPass := hashPassword(newPass)
	user := &model.User{
		Id: uid,
		Password: hashPass,
	}
	_, err := dao.db.Model(user).Set("password = ?password").WherePK().Update()
	if err != nil {
		fmt.Printf("change password error: %s\n", err.Error())
	}
	return err
}

func hashPassword(password string) string {
	data := []byte(password)
	hash := md5.Sum(data)
	hashPass := fmt.Sprintf("%x", hash)
	return hashPass
}