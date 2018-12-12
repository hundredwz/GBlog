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

type ArticleController struct {
	service.ContentService
}

// /article/delete /page/delete
func (ac *ArticleController) DelContent(c *gin.Context) {
	id := c.Query("cid")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": errors.New("wrong id"),
		})
		return
	}
	if cid, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": errors.New("wrong id"),
		})
		return
	} else {
		ac.ContentService.DeleteContent(&model.Content{Cid: cid})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "delete successfully",
	})
}

// /articles?p=*
func (ac *ArticleController) GetArticles(c *gin.Context) {
	p := c.DefaultQuery("p", "1")
	page := model.NewPage(p)
	params := map[string]interface{}{
		"type": "article",
	}
	if articles, err := ac.ContentService.GetArticles(page, params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"payload": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"payload": articles,
		})
	}
}

// /article?slug
func (ac *ArticleController) GetArticle(c *gin.Context) {
	id := c.Query("slug")
	articleId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(-1, gin.H{
			"status":  1,
			"payload": errors.New("article Id wrong"),
		})
	}
	params := make(map[string]interface{})
	params["type"] = "article"
	params["slug"] = articleId
	article, metas, err := ac.ContentService.GetArticle(params)
	if err != nil {
		c.JSON(-1, gin.H{
			"status":  1,
			"payload": err,
		})
		return
	}
	c.JSON(0, gin.H{
		"status":  0,
		"payload": []interface{}{article, metas},
	})
}

func (ac *ArticleController) UpdateContentStatus(c *gin.Context) {
	article := model.Content{}
	cid := c.PostForm("cid")
	if cid != "" {
		if v, err := strconv.Atoi(cid); err == nil {
			article.Cid = v
		}
	}
	status := c.PostForm("status")
	err := ac.ContentService.UpdateContentByMap(article, map[string]interface{}{"Status": status})
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

// /api/article/edit
func (ac *ArticleController) EditArticle(c *gin.Context) {
	article := model.Content{}
	cid := c.PostForm("cid")
	if cid != "" {
		if v, err := strconv.Atoi(cid); err == nil {
			article.Cid = v
		}
	}
	slug := c.PostForm("slug")
	article.Slug = slug
	title := c.PostForm("title")
	article.Title = title
	text := c.PostForm("text")
	article.Text = text
	created := c.PostForm("created")
	if created == "" || strings.Contains(created, "0001") {
		article.Created = time.Now()
	} else if v, err := time.Parse("2006-01-02T15:04", created); err != nil {
		article.Created = time.Now()
	} else {
		article.Created = v
	}
	article.Modified = time.Now()
	metas := make([]model.Meta, 0)
	categories := c.PostFormArray("category")
	for _, category := range categories {
		meta := model.Meta{Name: category, Type: "category"}
		metas = append(metas, meta)
	}
	tags := c.PostForm("tags")
	for _, tag := range strings.Split(tags, ";") {
		if tag == "" {
			continue
		}
		meta := model.Meta{Name: tag, Type: "tag"}
		metas = append(metas, meta)
	}
	allowComment := c.PostForm("allowComment")
	if allowComment != "" {
		if v, err := strconv.ParseBool(allowComment); err != nil {
			article.AllowComment = v
		}
	}
	allowPing := c.PostForm("allowPing")
	if allowPing != "" {
		if v, err := strconv.ParseBool(allowPing); err != nil {
			article.AllowPing = v
		}
	}
	allowFeed := c.PostForm("allowFeed")
	if allowFeed != "" {
		if v, err := strconv.ParseBool(allowFeed); err != nil {
			article.AllowFeed = v
		}
	}
	if status := c.PostForm("status"); status != "" {
		article.Status = status
	}
	article.Type = "article"
	if err := ac.ContentService.EditContent(article, metas); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": errors.New("edit article wrong"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "edit article successfully",
	})
}

// /api/page/edit
func (ac *ArticleController) EditPage(c *gin.Context) {
	page := model.Content{}
	cid := c.PostForm("cid")
	if cid != "" {
		if v, err := strconv.Atoi(cid); err == nil {
			page.Cid = v
		}
	}
	slug := c.PostForm("slug")
	page.Slug = slug
	title := c.PostForm("title")
	page.Title = title
	text := c.PostForm("text")
	page.Text = text
	created := c.PostForm("created")
	if created == "" || strings.Contains(created, "0001") {
		page.Created = time.Now()
	} else if v, err := time.Parse("2006-01-02T15:04", created); err != nil {
		page.Created = time.Now()
	} else {
		page.Created = v
	}
	page.Modified = time.Now()
	status := c.PostForm("status")
	page.Status = status
	allowComment := c.PostForm("allowComment")
	if allowComment != "" {
		if v, err := strconv.ParseBool(allowComment); err != nil {
			page.AllowComment = v
		}
	}
	allowPing := c.PostForm("allowPing")
	if allowPing != "" {
		if v, err := strconv.ParseBool(allowPing); err != nil {
			page.AllowPing = v
		}
	}
	allowFeed := c.PostForm("allowFeed")
	if allowFeed != "" {
		if v, err := strconv.ParseBool(allowFeed); err != nil {
			page.AllowFeed = v
		}
	}
	if status := c.PostForm("status"); status != "" {
		page.Status = status
	}
	page.Type = "page"
	if err := ac.ContentService.EditContent(page, nil); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": errors.New("edit page wrong"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "edit page successfully",
	})
}
