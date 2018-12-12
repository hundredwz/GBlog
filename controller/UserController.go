package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hundredwz/GBlog/model"
	"github.com/hundredwz/GBlog/service"
	"net/http"
)

type UserController struct {
	service.UserService
}

func (uc *UserController) Login(c *gin.Context) {

	name := c.PostForm("username")
	password := c.PostForm("password")
	user := model.User{
		Name:     name,
		Password: password,
	}
	result := uc.UserService.UserLogin(user)
	if !result {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": "login failed",
		})
		return
	}
	user = uc.UserService.GetUserByName(user.Name)
	token := model.NewToken(user.Uid, nil)
	c.SetCookie("user", token.Encode(), 2*24*60*3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "login successfully",
	})

}

func (uc *UserController) Logout(c *gin.Context) {
	c.SetCookie("user", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/")
}

func (uc *UserController) Register(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := model.User{
		Name:     name,
		Password: password,
	}
	if err := uc.UserService.AddUser(user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": "register failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  1,
		"payload": "register successfully",
	})
}
