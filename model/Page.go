package model

import (
	"github.com/hundredwz/GBlog/config"
	"strconv"
)

type Page struct {
	Index int
	Start int
	End   int
	Count int
}

func NewPage(i interface{}) Page {
	page := Page{}
	index := 1
	switch i.(type) {
	case int:
		index = i.(int)
	case string:
		if i, err := strconv.Atoi(i.(string)); err != nil {
			index = 1
		} else {
			index = i
		}
	default:
		index = 1
	}
	if index == 0 {
		index = 1
	}
	page.Index = index
	page.Start = (index - 1) * config.BlogArticleNumEachPage
	page.End = index*config.BlogArticleNumEachPage - 1
	page.Count = config.BlogArticleNumEachPage
	return page
}
