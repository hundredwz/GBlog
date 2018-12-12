package dao

import (
	"errors"
	"github.com/hundredwz/GBlog/model"
	"strings"
)

func (db *DataBase) CreateCommentTable() error {
	db.createTable(drop_gblog_comments)
	if err := db.createTable(gblog_comments); err != nil {
		return err
	}
	return nil
}

func (db *DataBase) GetCommentCount(params map[string]interface{}) int {

	query := "SELECT COUNT(*) FROM gblog_comments"
	args := make([]interface{}, 0)
	if params != nil && len(params) != 0 {
		query = query + " WHERE"
		for key, value := range params {
			query = query + " " + key + "=? AND"
			args = append(args, value)
		}
	}
	query = strings.TrimSuffix(query, " AND")
	results := db.get(query, &model.Count{}, args...)
	if len(results) == 0 {
		return 0
	}
	count := results[0].(model.Count)
	return count.Count
}

func (db *DataBase) GetComment(coid int) (model.Comment, error) {

	query := "SELECT * FROM gblog_comments WHERE Coid=?"
	results := db.get(query, &model.Comment{}, coid)
	if len(results) == 0 {
		return model.Comment{}, errors.New("no comment found")
	}
	return results[0].(model.Comment), nil

}

func (db *DataBase) GetComments(p model.Page, params map[string]interface{}) ([]model.Comment, error) {

	query := "SELECT * FROM gblog_comments"
	args := make([]interface{}, 0)
	if params != nil && len(params) != 0 {
		query = query + " WHERE"
		for key, value := range params {
			query = query + " " + key + "=? AND"
			args = append(args, value)
		}
	}
	query = strings.TrimSuffix(query, " AND")
	query = query + " ORDER BY Created desc LIMIT ?,?"
	args = append(args, p.Start, p.Count)
	results := db.get(query, &model.Comment{}, args...)
	comments := make([]model.Comment, 0)
	for _, result := range results {
		comment, ok := result.(model.Comment)
		if !ok {
			continue
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (db *DataBase) AddComment(comment model.Comment) (int, error) {

	sql := "INSERT INTO gblog_comments"
	insert, args := genInsert(comment)
	if insert == "" {
		return 0, errors.New("insert wrong")
	}
	sql = sql + insert
	row, err := db.modify(sql, args...)
	if err != nil {
		return 0, err
	}
	return row, nil
}

func (db *DataBase) UpdateComment(comment model.Comment) error {

	sql := "UPDATE gblog_comments "
	update, args := genUpdate(comment)
	if update != "" {
		sql = sql + update + " WHERE Coid=?"
		args = append(args, comment.Coid)
	}
	_, err := db.modify(sql, args...)
	return err
}

func (db *DataBase) UpdateCommentByMap(comment model.Comment, params map[string]interface{}) error {
	update := "UPDATE gblog_comments "
	args := make([]interface{}, 0)
	if params != nil && len(params) != 0 {
		update = update + " SET"
		for key, value := range params {
			update = update + " " + key + "=? ,"
			args = append(args, value)
		}
	}
	update = strings.TrimSuffix(update, " ,")
	update = update + " WHERE Coid=?"
	args = append(args, comment.Coid)
	_, err := db.modify(update, args...)
	return err
}

func (db *DataBase) DelComment(coid int) error {

	del := "DELETE FROM gblog_comments WHERE Coid=?"
	_, err := db.modify(del, coid)
	if err != nil {
		return err
	}
	return nil
}
