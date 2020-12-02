package service

import (
	xerrors "github.com/pkg/errors"
	"main/dao"
	"main/model"
)

type Service struct {
	dao *dao.Dao
}

func NewService() (s *Service, err error) {
	s.dao, err = dao.NewDao()
	if err != nil {
		return nil, xerrors.Wrap(err, "connect db failed")
	}
	return s, nil
}
func (s *Service) QueryUserById(id int) (*model.User, error) {
	user, err := s.dao.QueryUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
