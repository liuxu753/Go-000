package api

import (
	"fmt"
	"github.com/pkg/errors"
	"main/service"
)

func QueryUserById() {
	s, err := service.NewService()
	if err != nil {
		fmt.Printf("NewService failed,cause:%v,err:=%+v", errors.Unwrap(err), err)
		return
	}
	id := 999
	user, err := s.QueryUserById(id)
	if err != nil {
		fmt.Printf("query user failed,param：%d,cause:%v,err:=%+v", id, errors.Unwrap(err), err)
		return
	}
	fmt.Printf("user:%v", user)
	return
}
