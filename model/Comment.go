package model

import (
	"github.com/hundredwz/GBlog/config"
	"github.com/hundredwz/GBlog/util"
	"time"
)

type Comment struct {
	Coid     int       `orm:"Coid"json:"coid"`
	Cid      int       `orm:"Cid" json:"cid"`
	Created  time.Time `orm:"Created" json:"created"`
	Author   string    `orm:"Author" json:"author"`
	AuthorId int       `orm:"AuthorId" json:"authorId"`
	OwnerId  int       `orm:"OwnerId" json:"ownerId"`
	Mail     string    `orm:"Mail" json:"mail"`
	Url      string    `orm:"Url" json:"url"`
	Ip       string    `orm:"Ip" json:"ip"`
	Agent    string    `orm:"Agent" json:"agent"`
	Text     string    `orm:"Text" json:"text"`
	Type     string    `orm:"Type" json:"type"`
	Status   string    `orm:"Status" json:"status"`
	Parent   int       `orm:"Parent" json:"parent"`
}

func (c *Comment) TableName() string {
	return "gblog_comments"
}

func (c *Comment) ToHtml(article interface{}) map[string]interface{} {
	created := c.Created.Format(config.BlogCommentTimeFormat)
	avatarUrl := config.BlogCommentAvatarUrl + util.MD5(c.Mail)
	result := map[string]interface{}{
		"Coid":         c.Coid,
		"Cid":          c.Cid,
		"Created":      created,
		"Author":       c.Author,
		"AuthorId":     c.AuthorId,
		"AuthorAvatar": avatarUrl,
		"OwnerId":      c.OwnerId,
		"Mail":         c.Mail,
		"Url":          c.Url,
		"Ip":           c.Ip,
		"Agent":        c.Agent,
		"Text":         c.Text,
		"Type":         c.Type,
		"Status":       c.Status,
		"Parent":       c.Parent,
	}
	if article != nil {
		result["Article"] = article
	}
	return result
}
