package dao

import (
	"errors"
	"github.com/hundredwz/GBlog/model"
	"strconv"
	"strings"
)

func (db *DataBase) CreateContentTable() error {
	db.createTable(drop_gblog_contents)
	if err := db.createTable(gblog_contents); err != nil {
		return err
	}
	return nil
}

func (db *DataBase) GetContentCount(params map[string]interface{}) int {

	query := "SELECT COUNT(*) FROM gblog_contents"
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

func (db *DataBase) AddContent(article model.Content, metas []model.Meta) (int, error) {

	sql := "INSERT INTO gblog_contents"
	insert, args := genInsert(article)
	if insert == "" {
		return 0, errors.New("insert wrong")
	}
	sql = sql + insert
	row, err := db.modify(sql, args...)
	if err != nil {
		return 0, err
	}
	if article.Slug == "" {
		update := "UPDATE gblog_contents SET Slug=? WHERE Cid=?"
		db.modify(update, row, row)
	}
	article.Cid = row
	db.UpdateArticleMetas(article, metas)
	return row, nil
}

func (db *DataBase) UpdateContentByMap(article model.Content, params map[string]interface{}) error {

	update := "UPDATE gblog_contents "
	args := make([]interface{}, 0)
	if params != nil && len(params) != 0 {
		update = update + " SET"
		for key, value := range params {
			update = update + " " + key + "=? ,"
			args = append(args, value)
		}
	}
	update = strings.TrimSuffix(update, " ,")
	update = update + " WHERE Cid=?"
	args = append(args, article.Cid)
	_, err := db.modify(update, args...)
	return err
}

func (db *DataBase) UpdateContent(article model.Content, metas []model.Meta) error {

	sql := "UPDATE gblog_contents "
	update, args := genUpdate(article)
	if update != "" {
		sql = sql + update + " WHERE cid=?"
		args = append(args, article.Cid)
	}
	_, err := db.modify(sql, args...)
	db.UpdateArticleMetas(article, metas)
	return err
}

func (db *DataBase) UpdateContentBySql(article model.Content, sql string) error {

	update := "UPDATE gblog_contents "
	args := make([]interface{}, 0)
	if sql != "" {
		update = update + " SET " + sql
	}

	update = update + " WHERE Cid=?"
	args = append(args, article.Cid)
	_, err := db.modify(update, args...)
	return err
}

func (db *DataBase) UpdateArticleMetas(article model.Content, metas []model.Meta) {

	cmr := model.CMR{Cid: article.Cid}
	db.delCMR(cmr)
	query := "SELECT * FROM gblog_metas WHERE Name=? AND Type=?"
	for i := 0; i < len(metas); i++ {
		result := db.get(query, &model.Meta{}, metas[i].Name, metas[i].Type)
		if len(result) == 0 {
			metas[i].Count = 1
			row, _ := db.AddMeta(metas[i])
			metas[i].Mid = row
		} else {
			metas[i] = result[0].(model.Meta)
			metas[i].Count += 1
			db.UpdateMeta(metas[i])
		}
		cmr := model.CMR{Cid: article.Cid, Mid: metas[i].Mid}
		db.addCMR(cmr)
	}
}

func (db *DataBase) DeleteContent(article *model.Content) error {

	del := "DELETE FROM gblog_contents WHERE Cid=?"
	_, err := db.modify(del, article.Cid)
	if err != nil {
		return err
	}
	cmr := model.CMR{Cid: article.Cid}
	cmrs := db.getCMR(cmr)
	for _, c := range cmrs {
		if meta, err := db.GetMeta(map[string]interface{}{"Mid": c.Mid}); err == nil {
			meta.Count -= 1
			db.UpdateMeta(meta)
		}
	}
	db.delCMR(cmr)
	return nil
}

func (db *DataBase) GetArticle(params map[string]interface{}) (model.Content, []model.Meta, error) {

	query := "SELECT * FROM gblog_contents"
	args := make([]interface{}, 0)
	if params != nil && len(params) != 0 {
		query = query + " WHERE"
		for key, value := range params {
			query = query + " " + key + "=? AND"
			args = append(args, value)
		}
	}
	query = strings.TrimSuffix(query, " AND")
	results := db.get(query, &model.Content{}, args...)
	if len(results) == 0 {
		return model.Content{}, nil, nil
	}
	article := results[0].(model.Content)
	cmrs := db.getCMR(model.CMR{Cid: article.Cid})
	metas := make([]model.Meta, 0)
	for _, cmr := range cmrs {
		meta, err := db.GetMeta(map[string]interface{}{"Mid": cmr.Mid})
		if err != nil || meta.Mid == 0 {
			continue
		}
		metas = append(metas, meta)
	}
	return article, metas, nil
}

func (db *DataBase) GetArticles(p model.Page, params map[string]interface{}) ([]model.Content, error) {

	query := "SELECT * FROM gblog_contents"
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
	results := db.get(query, &model.Content{}, args...)
	articles := make([]model.Content, 0)
	for _, result := range results {
		article, ok := result.(model.Content)
		if !ok {
			continue
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (db *DataBase) GetMetaArticles(meta model.Meta, p model.Page, params map[string]interface{}) ([]model.Content, error) {

	query := "SELECT * FROM gblog_contents WHERE Cid in"
	args := make([]interface{}, 0)
	articles := make([]model.Content, 0)
	cmrs := db.getCMR(model.CMR{Mid: meta.Mid})
	cids := "("
	for _, cmr := range cmrs {
		str := strconv.Itoa(cmr.Cid)
		cids = cids + str + ","
	}
	cids = strings.TrimSuffix(cids, ",")
	cids = cids + ")"
	query = query + cids
	if params != nil && len(params) != 0 {
		query = query + " AND"
		for key, value := range params {
			query = query + " " + key + "=? AND"
			args = append(args, value)
		}
	}
	query = strings.TrimSuffix(query, " AND")
	query = query + " ORDER BY Created desc LIMIT ?,?"
	args = append(args, p.Start, p.Count)
	results := db.get(query, &model.Content{}, args...)
	for _, result := range results {
		article, ok := result.(model.Content)
		if !ok {
			continue
		}
		articles = append(articles, article)
	}
	return articles, nil
}
