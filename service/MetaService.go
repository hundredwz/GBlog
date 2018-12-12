package service

import (
	"github.com/hundredwz/GBlog/dao"
	"github.com/hundredwz/GBlog/model"
)

type MetaService struct {
	DB *dao.DataBase
}

func (ms *MetaService) CreateMetaTable() error {
	return ms.DB.CreateMetaTable()
}

func (ms *MetaService) GetCategoryCount() int {
	params := map[string]interface{}{"type": "category"}
	return ms.DB.GetMetaCount(params)
}

func (ms *MetaService) GetTagCount() int {
	params := map[string]interface{}{"type": "tag"}
	return ms.DB.GetMetaCount(params)
}

func (ms *MetaService) GetArticleMetas(article model.Content) ([]model.Meta, error) {
	return ms.DB.GetArticleMetas(article)
}

func (ms *MetaService) GetMeta(params map[string]interface{}) (model.Meta, error) {
	return ms.DB.GetMeta(params)
}

func (ms *MetaService) GetMetas(metaType string) ([]model.Meta, error) {
	return ms.DB.GetMetas(metaType)
}

func (ms *MetaService) EditMeta(meta model.Meta) error {
	if meta.Mid == 0 {
		return ms.AddMeta(meta)
	}
	return ms.UpdateMeta(meta)
}
func (ms *MetaService) AddMeta(meta model.Meta) error {
	_, err := ms.DB.AddMeta(meta)
	return err
}

func (ms *MetaService) UpdateMeta(meta model.Meta) error {
	return ms.DB.UpdateMeta(meta)
}

func (ms *MetaService) UpdateMetaByMap(meta model.Meta, params map[string]interface{}) error {
	return ms.DB.UpdateMetaByMap(meta, params)
}

func (ms *MetaService) DelMeta(mid int) error {
	meta := model.Meta{Mid: mid}
	return ms.DB.DelMeta(meta)
}
