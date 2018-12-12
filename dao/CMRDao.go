package dao

import (
	"github.com/hundredwz/GBlog/model"
)

func (db *DataBase) CreateCMRTable() error {
	db.createTable(drop_gblog_content_meta)
	if err := db.createTable(gblog_content_meta); err != nil {
		return err
	}
	return nil

}

func (db *DataBase) addCMR(cmr model.CMR) {
	if record := db.getCMR(cmr); len(record) != 0 {
		return
	}
	insert := "INSERT INTO gblog_content_meta(Cid,Mid) VALUES(?,?)"
	db.modify(insert, cmr.Cid, cmr.Mid)
	return
}

func (db *DataBase) delCMR(cmr model.CMR) {
	del := "DELETE FROM gblog_content_meta WHERE 1=1"
	args := make([]interface{}, 0)
	if cmr.Cid != 0 {
		del = del + " AND Cid=?"
		args = append(args, cmr.Cid)
	}
	if cmr.Mid != 0 {
		del = del + " AND Mid=?"
		args = append(args, cmr.Mid)
	}
	db.modify(del, args...)
	return
}

func (db *DataBase) getCMR(cmr model.CMR) []model.CMR {
	query := "SELECT * FROM gblog_content_meta WHERE 1=1 "
	args := make([]interface{}, 0)
	if cmr.Cid != 0 {
		query = query + " AND Cid=?"
		args = append(args, cmr.Cid)
	}
	if cmr.Mid != 0 {
		query = query + " AND Mid=?"
		args = append(args, cmr.Mid)
	}
	results := db.get(query, &model.CMR{}, args...)
	cmrs := make([]model.CMR, 0)
	for _, res := range results {
		v, ok := res.(model.CMR)
		if ok {
			cmrs = append(cmrs, v)
		}
	}
	return cmrs

}
