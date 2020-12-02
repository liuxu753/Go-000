package dao

import (
	"errors"
	xerrors "github.com/pkg/errors"
	"main/model"
)

type Dao struct {
}

func NewDao() (d *Dao, err error) {
	return &Dao{}, nil
}

func (d *Dao) QueryUserById(id int) (*model.User, error) {
	user, err := QueryRow("select * from user where `id`=?", id)
	if err != nil {
		return nil, xerrors.Wrapf(err, "get user failed id=%v", id)
	}
	return &user, nil
}

//模拟sql
var ErrNoRows = errors.New("sql ErrNoRows")

func QueryRow(query string, args ...interface{}) (user model.User, err error) {
	return model.User{}, ErrNoRows
}
