package service

import (
	"github.com/hundredwz/GBlog/dao"
	"github.com/hundredwz/GBlog/model"
)

type CommentService struct {
	DB *dao.DataBase
}

func (cs *CommentService) CreateCommentTable() error {
	return cs.CreateCommentTable()
}

func (cs *CommentService) GetCommentCount(params map[string]interface{}) int {
	if params == nil {
		params = map[string]interface{}{"type": "comment"}
	} else {
		params["type"] = "comment"
	}

	return cs.DB.GetCommentCount(params)
}

func (cs *CommentService) GetComment(coid int) (model.Comment, error) {
	return cs.DB.GetComment(coid)
}

func (cs *CommentService) GetComments(page model.Page, params map[string]interface{}) ([]model.Comment, error) {
	return cs.DB.GetComments(page, params)
}

func (cs *CommentService) EditComment(comment model.Comment) error {
	if comment.Coid == 0 {
		return cs.AddComment(comment)
	}
	return cs.UpdateComment(comment)
}

func (cs *CommentService) AddComment(comment model.Comment) error {
	_, err := cs.DB.AddComment(comment)
	if err != nil {
		return err
	}
	return cs.DB.UpdateContentBySql(model.Content{Cid: comment.Cid}, "CommentsNum=CommentsNum+1")
}

func (cs *CommentService) UpdateComment(comment model.Comment) error {
	return cs.DB.UpdateComment(comment)
}

func (cs *CommentService) UpdateCommentByMap(comment model.Comment, params map[string]interface{}) error {
	return cs.DB.UpdateCommentByMap(comment, params)
}

func (cs *CommentService) DelComment(coid int) error {
	comment, err := cs.GetComment(coid)
	if err != nil {
		return err
	}
	err = cs.DB.DelComment(coid)
	if err != nil {
		return err
	}
	return cs.DB.UpdateContentBySql(model.Content{Cid: comment.Cid}, "CommentsNum=CommentsNum-1")
}
