package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hundredwz/GBlog/model"
	"github.com/hundredwz/GBlog/service"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CommentController struct {
	service.CommentService
}

func (cc *CommentController) GetComment(c *gin.Context) {
	id := c.Query("coid")
	if coid, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": errors.New("comment id wrong"),
		})
		return
	} else {
		if comment, err := cc.CommentService.GetComment(coid); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  1,
				"payload": errors.New("comment id wrong"),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  0,
				"payload": comment,
			})
		}
	}
}

func (cc *CommentController) GetComments(c *gin.Context) {

}

func (cc *CommentController) EditComment(c *gin.Context) {
	comment := model.Comment{}
	coid := c.PostForm("coid")
	if coid != "" {
		if v, err := strconv.Atoi(coid); err == nil {
			comment.Coid = v
		}
	}
	cid := c.PostForm("cid")
	if v, err := strconv.Atoi(cid); err == nil {
		comment.Cid = v
	}
	created := c.PostForm("created")
	if created == "" || strings.Contains(created, "0001") {
		comment.Created = time.Now()
	} else if v, err := time.Parse("2006-01-02T15:04:05-07:00", created); err != nil {
		comment.Created = time.Now()
	} else {
		comment.Created = v
	}
	author := c.PostForm("author")
	comment.Author = author
	authorId := c.PostForm("authorId")
	if v, err := strconv.Atoi(authorId); err == nil {
		comment.AuthorId = v
	}
	mail := c.PostForm("mail")
	comment.Mail = mail
	url := c.PostForm("url")
	comment.Url = url
	ip := c.Request.RemoteAddr
	comment.Ip = ip
	agent := c.Request.UserAgent()
	comment.Agent = agent
	text := c.PostForm("text")
	comment.Text = text
	commentType := c.PostForm("type")
	comment.Type = commentType
	status := c.PostForm("status")
	comment.Status = status
	parent := c.PostForm("parent")
	if v, err := strconv.Atoi(parent); err == nil {
		comment.Parent = v
	}
	if err := cc.CommentService.EditComment(comment); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": errors.New("edit comment wrong"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "edit comment successfully",
	})
}

func (cc *CommentController) UpdateCommentStatus(c *gin.Context) {
	comment := model.Comment{}
	coid := c.PostForm("coid")
	if coid != "" {
		if v, err := strconv.Atoi(coid); err == nil {
			comment.Coid = v
		}
	}
	status := c.PostForm("status")
	err := cc.CommentService.UpdateCommentByMap(comment, map[string]interface{}{"Status": status})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": errors.New("update failed"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "update successfully",
	})
}

func (cc *CommentController) DelComment(c *gin.Context) {
	id := c.Query("coid")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": errors.New("wrong id"),
		})
		return
	}
	if coid, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": errors.New("wrong id"),
		})
		return
	} else {
		cc.CommentService.DelComment(coid)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "delete successfully",
	})

}
