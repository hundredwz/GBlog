package dao

import (
	"github.com/hundredwz/GBlog/model"
	"strconv"
	"strings"
)

func (db *DataBase) CreateMetaTable() error {
	db.createTable(drop_gblog_metas)
	if err := db.createTable(gblog_metas); err != nil {
		return err
	}
	return nil
}

func (db *DataBase) GetMetaCount(params map[string]interface{}) int {

	query := "SELECT COUNT(*) FROM gblog_metas"
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

func (db *DataBase) AddMeta(meta model.Meta) (int, error) {

	insert := "INSERT INTO gblog_metas(Name,Slug,Type,Description,Count,Order_,Parent) VALUES(?,?,?,?,?,?,?)"
	if meta.Slug == "" {
		meta.Slug = meta.Name
	}
	row, err := db.modify(insert, meta.Name, meta.Slug, meta.Type, meta.Description, meta.Count, meta.Order_, meta.Parent)
	if err != nil {
		return 0, err
	}
	return row, nil
}

func (db *DataBase) GetArticleMetas(article model.Content) ([]model.Meta, error) {

	query := "SELECT * FROM gblog_metas WHERE Mid in"
	metas := make([]model.Meta, 0)
	cmrs := db.getCMR(model.CMR{Cid: article.Cid})
	cids := "("
	for _, cmr := range cmrs {
		str := strconv.Itoa(cmr.Mid)
		cids = cids + str + ","
	}
	cids = strings.TrimSuffix(cids, ",")
	cids = cids + ")"
	query = query + cids
	results := db.get(query, &model.Meta{})
	for _, result := range results {
		meta, ok := result.(model.Meta)
		if !ok {
			continue
		}
		metas = append(metas, meta)
	}
	return metas, nil
}

func (db *DataBase) GetMeta(params map[string]interface{}) (model.Meta, error) {

	query := "SELECT * FROM gblog_metas"
	args := make([]interface{}, 0)
	if params != nil && len(params) != 0 {
		query = query + " WHERE"
		for key, value := range params {
			query = query + " " + key + "=? AND"
			args = append(args, value)
		}
	}
	query = strings.TrimSuffix(query, " AND")
	results := db.get(query, &model.Meta{}, args...)
	if len(results) == 0 {
		return model.Meta{}, nil
	}
	return results[0].(model.Meta), nil
}

func (db *DataBase) GetMetas(metaType string) ([]model.Meta, error) {

	metas := make([]model.Meta, 0)
	query := "SELECT * FROM gblog_metas WHERE Type=?"
	results := db.get(query, &model.Meta{}, metaType)
	for _, result := range results {
		meta, ok := result.(model.Meta)
		if !ok {
			continue
		}
		metas = append(metas, meta)
	}
	return metas, nil
}

func (db *DataBase) UpdateMeta(meta model.Meta) error {

	sql := "UPDATE gblog_metas "
	update, args := genUpdate(meta)
	if update != "" {
		sql = sql + update + " WHERE Mid=?"
		args = append(args, meta.Mid)
	}
	_, err := db.modify(sql, args...)
	return err
}

func (db *DataBase) UpdateMetaByMap(meta model.Meta, params map[string]interface{}) error {

	update := "UPDATE gblog_metas "
	args := make([]interface{}, 0)
	if params != nil && len(params) != 0 {
		update = update + " SET"
		for key, value := range params {
			update = update + " " + key + "=? ,"
			args = append(args, value)
		}
	}
	update = strings.TrimSuffix(update, " ,")
	update = update + " WHERE Mid=?"
	args = append(args, meta.Mid)
	_, err := db.modify(update, args...)
	return err
}

func (db *DataBase) DelMeta(meta model.Meta) error {

	del := "DELETE FROM gblog_metas WHERE Mid=?"
	_, err := db.modify(del, meta.Mid)
	if err != nil {
		return err
	}
	cmr := model.CMR{Mid: meta.Mid}
	db.delCMR(cmr)
	return nil
}
