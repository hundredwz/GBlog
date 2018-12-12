package service

import (
	"github.com/hundredwz/GBlog/dao"
	"github.com/hundredwz/GBlog/model"
)

type UserService struct {
	DB *dao.DataBase
}

func (us *UserService) AddUser(user model.User) error {
	_, err := us.DB.AddUser(user)
	return err
}

func (us *UserService) UpdateUserLogin(user model.User) error {
	return us.DB.UpdateUserLogin(user)
}

func (us *UserService) UpdateUser(user model.User) error {
	return us.DB.UpdateUser(user)
}

func (us *UserService) DelUser(user model.User) error {
	_, err := us.DB.DelUser(user)
	return err
}

func (us *UserService) GetUserByName(Name string) model.User {
	return us.DB.GetUserByName(Name)
}

func (us *UserService) GetUserById(id int) model.User {
	return us.DB.GetUserById(id)
}

func (us *UserService) UserLogin(user model.User) bool {
	return us.DB.UserLogin(user)
}
