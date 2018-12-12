package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hundredwz/GBlog/config"
	"github.com/hundredwz/GBlog/model"
	"github.com/hundredwz/GBlog/service"
	"net/http"
	"time"
)

type InstallController struct {
	service.InstallService
	service.UserService
}

func (ic *InstallController) InstallPage(c *gin.Context) {
	if !config.Installed {
		c.HTML(http.StatusOK, "install.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func (ic *InstallController) Install(c *gin.Context) {
	if config.Installed {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": "already installed",
		})
		return
	}
	blogName := c.PostForm("BlogName")
	config.BlogName = blogName
	blogDesc := c.PostForm("BlogDesc")
	config.BlogDesc = blogDesc
	blogKeywords := c.PostForm("BlogKeywords")
	config.BlogKeywords = blogKeywords
	dbName := c.PostForm("DBName")
	config.DBName = dbName
	dbUser := c.PostForm("DBUser")
	config.DBUser = dbUser
	dbPwd := c.PostForm("DBPwd")
	config.DBPwd = dbPwd
	username := c.PostForm("username")
	password := c.PostForm("password")
	err := ic.InstallService.Connection()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  2,
			"payload": "database connection error",
		})
		return
	}
	err = ic.InstallService.Install()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": "create database error",
		})
		return
	}
	ic.InstallService.Finish()
	user := model.User{
		Name:     username,
		Password: password,
		Created:  time.Now(),
		Logged:   time.Now(),
	}
	ic.UserService.AddUser(user)
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "blog install successfully",
	})
}

func (ic *InstallController) DBConnection(c *gin.Context) {
	dbName := c.PostForm("DBName")
	config.DBName = dbName
	dbUser := c.PostForm("DBUser")
	config.DBUser = dbUser
	dbPwd := c.PostForm("DBPwd")
	config.DBPwd = dbPwd
	err := ic.InstallService.Connection()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  2,
			"payload": "database connection error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "db fine",
	})
}
