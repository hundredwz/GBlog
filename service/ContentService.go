package service

import (
	"github.com/hundredwz/GBlog/config"
	"github.com/hundredwz/GBlog/dao"
	"github.com/hundredwz/GBlog/model"
	"strings"
)

type ContentService struct {
	DB *dao.DataBase
}

func (as *ContentService) CreateTable() error {
	return as.DB.CreateContentTable()
}

func (as *ContentService) GetContentCount() int {
	params := map[string]interface{}{"type": "article"}
	return as.DB.GetContentCount(params)
}

func (as *ContentService) GetPageCount() int {
	params := map[string]interface{}{"type": "page"}
	return as.DB.GetContentCount(params)
}

func (as *ContentService) EditContent(article model.Content, metas []model.Meta) error {
	if article.Cid == 0 {
		return as.AddContent(article, metas)
	}
	return as.UpdateContent(article, metas)
}

func (as *ContentService) AddContent(article model.Content, metas []model.Meta) error {
	_, err := as.DB.AddContent(article, metas)
	return err
}

func (as *ContentService) UpdateContent(article model.Content, metas []model.Meta) error {
	return as.DB.UpdateContent(article, metas)
}

func (as *ContentService) UpdateContentByMap(article model.Content, params map[string]interface{}) error {
	return as.DB.UpdateContentByMap(article, params)
}

func (as *ContentService) DeleteContent(article *model.Content) error {
	return as.DB.DeleteContent(article)
}

func (as *ContentService) GetContent(params map[string]interface{}) (model.Content, []model.Meta, error) {
	return as.DB.GetArticle(params)
}

func (as *ContentService) GetArticle(params map[string]interface{}) (model.Content, []model.Meta, error) {
	if params == nil {
		params = map[string]interface{}{
			"type": "article",
		}
	} else {
		params["type"] = "article"
	}
	return as.DB.GetArticle(params)
}

func (as *ContentService) GetPage(params map[string]interface{}) (model.Content, []model.Meta, error) {
	if params == nil {
		params = map[string]interface{}{
			"type": "page",
		}
	} else {
		params["type"] = "page"
	}
	return as.DB.GetArticle(params)
}

func (as *ContentService) GetArticles(p model.Page, params map[string]interface{}) ([]model.Content, error) {
	if params == nil {
		params = map[string]interface{}{
			"type": "article",
		}
	} else {
		params["type"] = "article"
	}

	articles, err := as.DB.GetArticles(p, params)
	if !config.BlogArticleSub {
		return articles, err
	}
	for i, l := 0, len(articles); i < l; i++ {
		var text string
		if strings.Contains(articles[i].Text, "<--more-->") {
			text = strings.SplitN(articles[i].Text, "<--more-->", 2)[0]
		} else {
			text = strings.SplitN(articles[i].Text, "\n", 2)[0]
		}
		articles[i].Text = text
	}
	return articles, err
}

func (as *ContentService) GetPages(p model.Page, params map[string]interface{}) ([]model.Content, error) {
	if params == nil {
		params = map[string]interface{}{
			"type": "page",
		}
	} else {
		params["type"] = "page"
	}
	articles, err := as.DB.GetArticles(p, params)
	if !config.BlogArticleSub {
		return articles, err
	}
	for i, l := 0, len(articles); i < l; i++ {
		var text string
		if strings.Contains(articles[i].Text, "<--more-->") {
			text = strings.SplitN(articles[i].Text, "<--more-->", 2)[0]
		} else {
			text = strings.SplitN(articles[i].Text, "\n", 2)[0]
		}
		articles[i].Text = text
	}
	return articles, err
}

func (as *ContentService) GetMetaArticles(meta model.Meta, p model.Page, params map[string]interface{}) ([]model.Content, error) {
	articles, err := as.DB.GetMetaArticles(meta, p, params)
	if !config.BlogArticleSub {
		return articles, err
	}
	for i, l := 0, len(articles); i < l; i++ {
		var text string
		if strings.Contains(articles[i].Text, "<--more-->") {
			text = strings.SplitN(articles[i].Text, "<--more-->", 2)[0]
		} else {
			text = strings.SplitN(articles[i].Text, "\n", 2)[0]
		}
		articles[i].Text = text
	}
	return articles, err
}
