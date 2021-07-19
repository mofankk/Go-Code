package service

import "errors"

type IUserService interface {
	GetName(id int) string
	DelUser(id int) error
}

type UserService struct {}

func (UserService) GetName(id int) string {
	if id == 101 {
		return "Mofan"
	}
	return "Guest"
}

func (UserService) DelUser(id int) error {
	if id == 102 {
		return nil
	} else {
		return errors.New("删除失败")
	}
}