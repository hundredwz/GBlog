package model

import (
	"github.com/hundredwz/GBlog/config"
	"github.com/russross/blackfriday"
	"html/template"
	"strings"
	"time"
)

type Content struct {
	Cid          int       `orm:"Cid" json:"Cid"`
	Title        string    `orm:"Title" json:"Title"`
	Slug         string    `orm:"Slug" json:"slug"`
	Created      time.Time `orm:"Created" json:"Created"`
	Modified     time.Time `orm:"Modified" json:"Modified"`
	Text         string    `orm:"Text" json:"Text"`
	Order_       int       `orm:"Order_" json:"Order"`
	AuthorId     int       `orm:"AuthorId" json:"AuthorId"`
	Template     string    `orm:"Template" json:"Template"`
	Type         string    `orm:"Type" json:"Type"`
	Status       string    `orm:"Status" json:"Status"`
	Password     string    `orm:"Password" json:"Password"`
	CommentsNum  int       `orm:"CommentsNum" json:"CommentsNum"`
	AllowComment bool      `orm:"AllowComment" json:"AllowComment"`
	AllowPing    bool      `orm:"AllowPing" json:"AllowPing"`
	AllowFeed    bool      `orm:"AllowFeed" json:"AllowFeed"`
	Parent       int       `orm:"Parent" json:"Parent"`
}

func (c *Content) ToMarkdown(metas []Meta) map[string]interface{} {
	created := c.Created.Format(config.BlogArticleTimeFormat)
	modified := c.Modified.Format(config.BlogArticleTimeFormat)
	tags := make([]Meta, 0)
	category := Meta{}
	if metas != nil {
		for _, meta := range metas {
			if meta.Type == "category" {
				category = meta
			} else if meta.Type == "tag" {
				tags = append(tags, meta)
			}
		}
	}

	t := strings.Replace(c.Text, "\r\n", "\n", -1)
	text := template.HTML(blackfriday.Run([]byte(t)))
	return map[string]interface{}{
		"Cid":          c.Cid,
		"Title":        c.Title,
		"Slug":         c.Slug,
		"Created":      created,
		"Modified":     modified,
		"Text":         text,
		"Order":        c.Order_,
		"AuthorId":     c.AuthorId,
		"Template":     c.Template,
		"Type":         c.Type,
		"Status":       c.Status,
		"Password":     c.Password,
		"CommentsNum":  c.CommentsNum,
		"AllowComment": c.AllowComment,
		"AllowPing":    c.AllowPing,
		"AllowFeed":    c.AllowFeed,
		"Parent":       c.Parent,
		"Tags":         tags,
		"Category":     category,
	}
}
