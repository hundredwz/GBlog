package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hundredwz/GBlog/config"
	"github.com/hundredwz/GBlog/model"
	"github.com/hundredwz/GBlog/service"
	"net/http"
	"strconv"
)

type SettingController struct {
	service.UserService
}

func (sc *SettingController) BlogSetting(c *gin.Context) {
	blogName := c.PostForm("BlogName")
	config.BlogName = blogName
	blogDesc := c.PostForm("BlogDesc")
	config.BlogDesc = blogDesc
	blogKeywords := c.PostForm("BlogKeywords")
	config.BlogKeywords = blogKeywords
	blogCommentTimeFormat := c.PostForm("BlogCommentTimeFormat")
	config.BlogCommentTimeFormat = blogCommentTimeFormat
	blogCommentListNum := c.PostForm("BlogCommentListNum")
	if v, err := strconv.Atoi(blogCommentListNum); err == nil {
		config.BlogCommentListNum = v
	}
	blogCommentAvatarUrl := c.PostForm("BlogCommentAvatarUrl")
	config.BlogCommentAvatarUrl = blogCommentAvatarUrl
	blogArticleTimeFormat := c.PostForm("BlogArticleTimeFormat")
	config.BlogArticleTimeFormat = blogArticleTimeFormat
	blogArticleNumEachPage := c.PostForm("BlogArticleNumEachPage")
	if v, err := strconv.Atoi(blogArticleNumEachPage); err == nil {
		config.BlogArticleNumEachPage = v
	}
	blogArticleSub := c.PostForm("BlogArticleSub")
	if v, err := strconv.ParseBool(blogArticleSub); err == nil {
		config.BlogArticleSub = v
	}
	if err := config.UpdateConfig(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": "update failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "update successfully",
	})

}

func (sc *SettingController) UserSetting(c *gin.Context) {
	user := model.User{}
	uid := c.PostForm("Uid")
	fmt.Println(uid)
	if v, err := strconv.Atoi(uid); err == nil {
		user.Uid = v
	}
	name := c.PostForm("Name")
	user.Name = name
	password := c.PostForm("Password")
	user.Password = password
	mail := c.PostForm("Mail")
	user.Mail = mail
	url := c.PostForm("Url")
	user.Url = url
	if err := sc.UserService.UpdateUser(user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": "update failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "update successfully",
	})
}
