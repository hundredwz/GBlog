package service

import (
	"github.com/hundredwz/GBlog/config"
	"github.com/hundredwz/GBlog/dao"
)

type InstallService struct {
	DB *dao.DataBase
}

func (is *InstallService) Connection() error {
	return is.DB.Connection()
}

func (is *InstallService) Install() error {
	err := is.DB.CreateContentTable()
	if err != nil {
		return err
	}
	err = is.DB.CreateMetaTable()
	if err != nil {
		return err
	}
	err = is.DB.CreateCMRTable()
	if err != nil {
		return err
	}
	err = is.DB.CreateCommentTable()
	if err != nil {
		return err
	}
	err = is.DB.CreateUserTable()
	if err != nil {
		return err
	}
	return nil
}
func (is *InstallService) Finish() {
	config.Installed = true
	config.UpdateConfig()
}
