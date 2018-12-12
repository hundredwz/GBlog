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

//admin

type AdminController struct {
	service.MetaService
	service.ContentService
	service.CommentService
	service.UserService
}

// /admin/index
func (ac *AdminController) Index(c *gin.Context) {
	var user model.User
	if uid, ok := c.Get("userId"); ok {
		id := uid.(int)
		user = ac.UserService.GetUserById(id)
	}

	articleCount, pageCount := ac.ContentService.GetContentCount(), ac.ContentService.GetPageCount()
	categoryCount, tagCount := ac.MetaService.GetCategoryCount(), ac.MetaService.GetTagCount()
	commentCount := ac.CommentService.GetCommentCount(nil)

	articles, _ := ac.ContentService.GetArticles(model.NewPage(1), nil)
	pages, _ := ac.ContentService.GetPages(model.NewPage(1), nil)
	comments, _ := ac.CommentService.GetComments(model.NewPage(1), nil)
	h := gin.H{
		"User":          user,
		"ArticleCount":  articleCount,
		"PageCount":     pageCount,
		"CategoryCount": categoryCount,
		"TagCount":      tagCount,
		"CommentCount":  commentCount,
		"Articles":      articles,
		"Pages":         pages,
		"Comments":      comments,
	}
	c.HTML(http.StatusOK, "admin-index.html", h)
}

func (ac *AdminController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-login.html", nil)
}

// /admin/article/edit
func (ac *AdminController) ArticleEdit(c *gin.Context) {
	articleSlug := c.Query("slug")
	var user model.User
	if uid, ok := c.Get("userId"); ok {
		id := uid.(int)
		user = ac.UserService.GetUserById(id)
	}
	article := model.Content{}
	metas := make([]model.Meta, 0)
	categories, _ := ac.MetaService.GetMetas("category")
	h := gin.H{
		"User":       user,
		"Categories": categories,
	}
	if articleSlug != "" {
		article, metas, _ = ac.ContentService.GetArticle(map[string]interface{}{"type": "post", "slug": articleSlug})
		if article.Cid != 0 {
			tags := make([]model.Meta, 0)
			articleCategories := make([]int, 0)
			for _, meta := range metas {
				if meta.Type == "tag" {
					tags = append(tags, meta)
				} else if meta.Type == "category" {
					articleCategories = append(articleCategories, meta.Mid)
				}
			}
			h["Tags"] = tags
			h["ArticleCatIds"] = articleCategories
		}
	}
	h["Article"] = article
	c.HTML(http.StatusOK, "admin-article-edit.html", h)
}

// /admin/article/list
func (ac *AdminController) ArticleList(c *gin.Context) {
	p := c.DefaultQuery("p", "1")
	page := model.NewPage(p)
	var user model.User
	if uid, ok := c.Get("userId"); ok {
		id := uid.(int)
		user = ac.UserService.GetUserById(id)
	}
	categories, _ := ac.MetaService.GetMetas("category")
	articleCount := ac.ContentService.GetContentCount()
	var count int
	if articleCount%config.BlogArticleNumEachPage == 0 {
		count = articleCount / config.BlogArticleNumEachPage
	} else {
		count = articleCount/config.BlogArticleNumEachPage + 1
	}
	allPages := make([]model.Page, count)

	for i := 0; i < count; i++ {
		allPages[i] = model.NewPage(i + 1)
	}
	h := gin.H{
		"User":       user,
		"Categories": categories,
		"CurrPage":   page,
		"AllPages":   allPages,
	}
	if articles, err := ac.ContentService.GetArticles(page, nil); err == nil {
		results := make([]map[string]interface{}, 0)
		for _, article := range articles {
			metas, _ := ac.MetaService.GetArticleMetas(article)
			results = append(results, article.ToMarkdown(metas))
		}
		h["Articles"] = results
	}
	c.HTML(http.StatusOK, "admin-article-list.html", h)
}

// /admin/page/edit
func (ac *AdminController) PageEdit(c *gin.Context) {
	pageSlug := c.Query("slug")
	var user model.User
	if uid, ok := c.Get("userId"); ok {
		id := uid.(int)
		user = ac.UserService.GetUserById(id)
	}
	page := model.Content{}
	var metas []model.Meta
	categories, _ := ac.MetaService.GetMetas("category")
	h := gin.H{
		"User":       user,
		"Categories": categories,
	}
	if pageSlug != "" {
		page, metas, _ = ac.ContentService.GetPage(map[string]interface{}{"type": "page", "slug": pageSlug})
	}
	h["Page"] = page
	h["Metas"] = metas
	c.HTML(http.StatusOK, "admin-page-edit.html", h)
}

// /admin/page/list
func (ac *AdminController) PageList(c *gin.Context) {
	p := c.DefaultQuery("p", "1")
	page := model.NewPage(p)
	var user model.User
	if uid, ok := c.Get("userId"); ok {
		id := uid.(int)
		user = ac.UserService.GetUserById(id)
	}
	categories, _ := ac.MetaService.GetMetas("category")
	pageCount := ac.ContentService.GetPageCount()
	var count int
	if pageCount%config.BlogArticleNumEachPage == 0 {
		count = pageCount / config.BlogArticleNumEachPage
	} else {
		count = pageCount/config.BlogArticleNumEachPage + 1
	}
	allPages := make([]model.Page, count)

	for i := 0; i < count; i++ {
		allPages[i] = model.NewPage(i + 1)
	}
	h := gin.H{
		"User":       user,
		"Categories": categories,
		"CurrPage":   page,
		"AllPages":   allPages,
	}
	if pages, err := ac.ContentService.GetPages(page, nil); err == nil {
		results := make([]map[string]interface{}, 0)
		for _, article := range pages {
			metas, _ := ac.MetaService.GetArticleMetas(article)
			results = append(results, article.ToMarkdown(metas))
		}
		h["Pages"] = results
	}
	c.HTML(http.StatusOK, "admin-page-list.html", h)
}

// /admin/category/edit
func (ac *AdminController) CategoryEdit(c *gin.Context) {
	metaSlug := c.Query("slug")
	var user model.User
	if uid, ok := c.Get("userId"); ok {
		id := uid.(int)
		user = ac.UserService.GetUserById(id)
	}
	meta := model.Meta{}
	categories, _ := ac.MetaService.GetMetas("category")
	h := gin.H{
		"User":       user,
		"Categories": categories,
		"MetaType":   "category",
	}
	if metaSlug != "" {
		meta, _ = ac.MetaService.GetMeta(map[string]interface{}{"type": "category", "slug": metaSlug})
	}
	h["Meta"] = meta
	c.HTML(http.StatusOK, "admin-category-edit.html", h)
}

func (ac *AdminController) CategoryList(c *gin.Context) {
	var user model.User
	if uid, ok := c.Get("userId"); ok {
		id := uid.(int)
		user = ac.UserService.GetUserById(id)
	}
	categories, _ := ac.MetaService.GetMetas("category")
	h := gin.H{
		"User":       user,
		"Categories": categories,
	}
	if categories, err := ac.MetaService.GetMetas("category"); err == nil {
		h["CategoryList"] = categories
	}
	c.HTML(http.StatusOK, "admin-category-list.html", h)
}

// /admin/tag/edit
func (ac *AdminController) TagEdit(c *gin.Context) {
	metaSlug := c.Query("slug")
	var user model.User
	if uid, ok := c.Get("userId"); ok {
		id := uid.(int)
		user = ac.UserService.GetUserById(id)
	}
	meta := model.Meta{}
	categories, _ := ac.MetaService.GetMetas("category")
	h := gin.H{
		"User":       user,
		"Categories": categories,
		"MetaType":   "tag",
	}
	if metaSlug != "" {
		meta, _ = ac.MetaService.GetMeta(map[string]interface{}{"type": "tag", "slug": metaSlug})
	}
	h["Meta"] = meta
	c.HTML(http.StatusOK, "admin-category-edit.html", h)
}

func (ac *AdminController) TagList(c *gin.Context) {
	var user model.User
	if uid, ok := c.Get("userId"); ok {
		id := uid.(int)
		user = ac.UserService.GetUserById(id)
	}
	categories, _ := ac.MetaService.GetMetas("category")
	h := gin.H{
		"User":       user,
		"Categories": categories,
	}
	if categories, err := ac.MetaService.GetMetas("tag"); err == nil {
		h["Tags"] = categories
	}
	c.HTML(http.StatusOK, "admin-tag-list.html", h)
}

func (ac *AdminController) CommentList(c *gin.Context) {
	p := c.DefaultQuery("p", "1")
	params := make(map[string]interface{})
	if status := c.Query("s"); status != "" {
		params["status"] = status
	}
	if id := c.Query("cid"); id != "" {
		if cid, err := strconv.Atoi(id); err == nil {
			params["cid"] = cid
		}
	}
	page := model.NewPage(p)
	var user model.User
	if uid, ok := c.Get("userId"); ok {
		id := uid.(int)
		user = ac.UserService.GetUserById(id)
	}
	categories, _ := ac.MetaService.GetMetas("category")
	commentCount := ac.CommentService.GetCommentCount(params)
	var count int
	if commentCount%config.BlogCommentListNum == 0 {
		count = commentCount / config.BlogCommentListNum
	} else {
		count = commentCount/config.BlogCommentListNum + 1
	}
	allPages := make([]model.Page, count)

	for i := 0; i < count; i++ {
		allPages[i] = model.NewPage(i + 1)
	}
	h := gin.H{
		"User":       user,
		"Categories": categories,
		"CurrPage":   page,
		"AllPages":   allPages,
	}
	if comments, err := ac.CommentService.GetComments(page, params); err == nil {
		results := make([]map[string]interface{}, 0)
		for _, comment := range comments {
			article, _, _ := ac.ContentService.GetContent(map[string]interface{}{"Cid": comment.Cid})
			results = append(results, comment.ToHtml(article))
		}
		h["Comments"] = results
	}
	c.HTML(http.StatusOK, "admin-comment-list.html", h)
}

func (ac *AdminController) BlogSetting(c *gin.Context) {
	var user model.User
	if uid, ok := c.Get("userId"); ok {
		id := uid.(int)
		user = ac.UserService.GetUserById(id)
	}
	categories, _ := ac.MetaService.GetMetas("category")
	h := gin.H{
		"User":                   user,
		"Categories":             categories,
		"BlogName":               config.BlogName,
		"BlogKeywords":           config.BlogKeywords,
		"BlogDesc":               config.BlogDesc,
		"BlogCommentTimeFormat":  config.BlogCommentTimeFormat,
		"BlogCommentListNum":     config.BlogCommentListNum,
		"BlogCommentAvatarUrl":   config.BlogCommentAvatarUrl,
		"BlogArticleTimeFormat":  config.BlogArticleTimeFormat,
		"BlogArticleNumEachPage": config.BlogArticleNumEachPage,
		"BlogArticleSub":         config.BlogArticleSub,
	}
	c.HTML(http.StatusOK, "admin-setting-blog.html", h)
}

func (ac *AdminController) UserSetting(c *gin.Context) {
	var user model.User
	if uid, ok := c.Get("userId"); ok {
		id := uid.(int)
		user = ac.UserService.GetUserById(id)
	}
	fmt.Println(user)
	categories, _ := ac.MetaService.GetMetas("category")
	h := gin.H{
		"User":       user,
		"Categories": categories,
	}
	c.HTML(http.StatusOK, "admin-setting-user.html", h)
}
