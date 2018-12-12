package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hundredwz/GBlog/model"
	"github.com/hundredwz/GBlog/service"
	"net/http"
	"strconv"
)

type MetaController struct {
	service.MetaService
	service.ContentService
}

// /category/delete /tag/delete
func (mc *MetaController) DelMeta(c *gin.Context) {
	id := c.Query("mid")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": errors.New("wrong category id"),
		})
		return
	}
	if mid, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": errors.New("wrong category id"),
		})
		return
	} else {
		mc.MetaService.DelMeta(mid)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "delete category successfully",
	})
}

func (mc *MetaController) UpdateMetaStatus(c *gin.Context) {
	meta := model.Meta{}
	mid := c.PostForm("mid")
	if mid != "" {
		if v, err := strconv.Atoi(mid); err == nil {
			meta.Mid = v
		}
	}
	status := c.PostForm("status")
	err := mc.MetaService.UpdateMetaByMap(meta, map[string]interface{}{"Status": status})
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

// /category/:name?p=
func (mc *MetaController) Category(c *gin.Context) {
	cat := c.Param("name")
	p := c.DefaultQuery("p", "1")
	page := model.NewPage(p)
	params := make(map[string]interface{})
	params["type"] = "category"
	params["name"] = cat
	var meta model.Meta
	meta, err := mc.MetaService.GetMeta(params)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": err,
		})
		return
	}
	articleParams := map[string]interface{}{
		"Status": "publish",
	}
	articles, err := mc.ContentService.GetMetaArticles(meta, page, articleParams)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": articles,
	})
}

// /category/:name/info
func (mc *MetaController) CategoryInfo(c *gin.Context) {
	slug := c.Param("name")
	params := make(map[string]interface{})
	params["type"] = "category"
	params["slug"] = slug
	var meta model.Meta
	meta, err := mc.MetaService.GetMeta(params)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": meta,
	})
}

// /category/list
func (mc *MetaController) CategoryList(c *gin.Context) {
	metas, err := mc.MetaService.GetMetas("category")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": metas,
	})
}

// /category/edit
func (mc *MetaController) EditCategory(c *gin.Context) {
	meta := model.Meta{}
	mid := c.PostForm("mid")
	if mid != "" {
		if v, err := strconv.Atoi(mid); err == nil {
			meta.Mid = v
		}
	}
	meta.Type = "category"
	slug := c.PostForm("slug")
	meta.Slug = slug
	name := c.PostForm("name")
	meta.Name = name
	desc := c.PostForm("description")
	meta.Description = desc

	if err := mc.MetaService.EditMeta(meta); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": errors.New("edit meta wrong"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "edit meta successfully",
	})
}

// /tag/:name?p=
func (mc *MetaController) Tag(c *gin.Context) {
	slug := c.Param("name")
	p := c.DefaultQuery("p", "1")
	page := model.NewPage(p)
	params := make(map[string]interface{})
	params["type"] = "tag"
	params["slug"] = slug
	meta, err := mc.MetaService.GetMeta(params)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": "get tag failed",
		})
		return
	}
	articleParams := map[string]interface{}{
		"Status": "publish",
	}
	articles, err := mc.ContentService.GetMetaArticles(meta, page, articleParams)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": "get article failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": articles,
	})
}

// /tag
func (mc *MetaController) TagInfo(c *gin.Context) {
	slug := c.Query("slug")
	params := make(map[string]interface{})
	params["type"] = "tag"
	params["slug"] = slug
	var meta model.Meta
	meta, err := mc.MetaService.GetMeta(params)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": "edit page successfully",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": meta,
	})
}

// /tag/list
func (mc *MetaController) TagList(c *gin.Context) {
	metas, err := mc.MetaService.GetMetas("tag")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": metas,
	})
}

func (mc *MetaController) EditTag(c *gin.Context) {
	meta := model.Meta{}
	mid := c.PostForm("mid")
	if mid != "" {
		if v, err := strconv.Atoi(mid); err == nil {
			meta.Mid = v
		}
	}
	meta.Type = "tag"
	slug := c.PostForm("slug")
	meta.Slug = slug
	name := c.PostForm("name")
	meta.Name = name
	desc := c.PostForm("description")
	meta.Description = desc

	if err := mc.MetaService.EditMeta(meta); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"payload": errors.New("edit meta wrong"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"payload": "edit meta successfully",
	})
}
