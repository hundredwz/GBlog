package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hundredwz/GBlog/config"
	"github.com/hundredwz/GBlog/model"
	"github.com/hundredwz/GBlog/service"
	"net/http"
)

type PageController struct {
	service.ContentService
	service.MetaService
	service.CommentService
}

// index?p=
func (pc *PageController) Index(c *gin.Context) {
	p := c.DefaultQuery("p", "1")
	page := model.NewPage(p)
	nextPage := page.Index + 1
	prevPage := page.Index - 1
	categories, _ := pc.MetaService.GetMetas("category")
	h := gin.H{
		"BlogName":     config.BlogName,
		"BlogDesc":     config.BlogDesc,
		"BlogKeywords": config.BlogKeywords,
		"Categories":   categories,
		"NextPage":     nextPage,
		"PrevPage":     prevPage,
	}
	pages, _ := pc.ContentService.GetPages(model.NewPage(1), map[string]interface{}{"Status": "publish"})
	h["Pages"] = pages
	if articles, err := pc.ContentService.GetArticles(page, map[string]interface{}{"Status": "publish"}); err == nil {
		results := make([]map[string]interface{}, 0)
		for _, article := range articles {
			results = append(results, article.ToMarkdown(nil))
		}
		h["Articles"] = results
	}
	c.HTML(http.StatusOK, "index.html", h)
}

// /article?slug=
func (pc *PageController) Article(c *gin.Context) {
	slug := c.Query("slug")
	cp := c.DefaultQuery("cp", "1")

	pages, _ := pc.ContentService.GetPages(model.NewPage(1), map[string]interface{}{"Status": "publish"})

	categories, _ := pc.MetaService.GetMetas("category")

	h := gin.H{
		"BlogName":     config.BlogName,
		"BlogDesc":     config.BlogDesc,
		"BlogKeywords": config.BlogKeywords,
		"Categories":   categories,
		"Pages":        pages,
	}
	article, metas, err := pc.ContentService.GetArticle(map[string]interface{}{"slug": slug})
	if err != nil || article.Status != "publish" {
		c.HTML(http.StatusOK, "article.html", h)
		return
	} else {
		h["Article"] = article.ToMarkdown(metas)
	}
	if comments, err := pc.CommentService.GetComments(model.NewPage(cp), map[string]interface{}{
		"cid":    article.Cid,
		"status": "approved",
	}); err == nil {
		results := make([]map[string]interface{}, 0)
		for _, comment := range comments {
			results = append(results, comment.ToHtml(nil))
		}
		h["Comments"] = results
	}
	c.HTML(http.StatusOK, "article.html", h)
}

// /category/:slug?p=
func (pc *PageController) Category(c *gin.Context) {
	p := c.DefaultQuery("p", "1")
	slug := c.Param("slug")
	page := model.NewPage(p)
	nextPage := page.Index + 1
	prevPage := page.Index - 1
	categories, _ := pc.MetaService.GetMetas("category")
	pages, _ := pc.ContentService.GetPages(model.NewPage(1), map[string]interface{}{"Status": "publish"})
	h := gin.H{
		"BlogName":     config.BlogName,
		"BlogDesc":     config.BlogDesc,
		"BlogKeywords": config.BlogKeywords,
		"Categories":   categories,
		"NextPage":     nextPage,
		"PrevPage":     prevPage,
		"Pages":        pages,
	}
	metaParams := map[string]interface{}{
		"type": "category",
		"slug": slug,
	}
	articleParams := map[string]interface{}{
		"Status": "publish",
	}
	if meta, err := pc.MetaService.GetMeta(metaParams); err == nil {
		if articles, err := pc.ContentService.GetMetaArticles(meta, page, articleParams); err == nil {
			results := make([]map[string]interface{}, 0)
			for _, article := range articles {
				results = append(results, article.ToMarkdown(nil))
			}
			h["Articles"] = results
		}
		h["Meta"] = meta
	}
	c.HTML(http.StatusOK, "meta.html", h)
}

// /tag
func (pc *PageController) Tag(c *gin.Context) {
	p := c.DefaultQuery("p", "1")
	slug := c.Param("slug")
	page := model.NewPage(p)
	nextPage := page.Index + 1
	prevPage := page.Index - 1
	categories, _ := pc.MetaService.GetMetas("category")
	pages, _ := pc.ContentService.GetPages(model.NewPage(1), map[string]interface{}{"Status": "publish"})
	h := gin.H{
		"BlogName":     config.BlogName,
		"BlogDesc":     config.BlogDesc,
		"BlogKeywords": config.BlogKeywords,
		"Categories":   categories,
		"Pages":        pages,
		"NextPage":     nextPage,
		"PrevPage":     prevPage,
	}
	metaParams := map[string]interface{}{
		"type": "tag",
		"slug": slug,
	}
	articleParams := map[string]interface{}{
		"Status": "publish",
	}
	if meta, err := pc.MetaService.GetMeta(metaParams); err == nil {
		if articles, err := pc.ContentService.GetMetaArticles(meta, page, articleParams); err == nil {
			results := make([]map[string]interface{}, 0)
			for _, article := range articles {
				results = append(results, article.ToMarkdown(nil))
			}
			h["Articles"] = results
		}
		h["Meta"] = meta
	}
	c.HTML(http.StatusOK, "meta.html", h)
}

// /page?slug=
func (pc *PageController) Page(c *gin.Context) {
	slug := c.Query("slug")
	cp := c.DefaultQuery("cp", "1")
	article, metas, err := pc.ContentService.GetPage(map[string]interface{}{"slug": slug})
	if err != nil {
		return
	}
	categories, _ := pc.MetaService.GetMetas("category")
	pages, _ := pc.ContentService.GetPages(model.NewPage(1), map[string]interface{}{"Status": "publish"})
	h := gin.H{
		"BlogName":     config.BlogName,
		"BlogDesc":     config.BlogDesc,
		"BlogKeywords": config.BlogKeywords,
		"Categories":   categories,
		"Pages":        pages,
		"Article":      article.ToMarkdown(metas),
	}
	if comments, err := pc.CommentService.GetComments(model.NewPage(cp), map[string]interface{}{
		"Cid":    article.Cid,
		"status": "approved",
	}); err == nil {
		results := make([]map[string]interface{}, 0)
		for _, comment := range comments {
			results = append(results, comment.ToHtml(nil))
		}
		h["Comments"] = results

	}
	c.HTML(http.StatusOK, "article.html", h)
}
