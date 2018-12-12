package model

type Meta struct {
	Mid         int    `orm:"Mid" json:"mid"`
	Name        string `orm:"Name" json:"name"`
	Slug        string `orm:"Slug" json:"slug"`
	Type        string `orm:"Type" json:"type"`
	Description string `orm:"Description" json:"description"`
	Count       int    `orm:"Count" json:"count"`
	Order_      int    `orm:"Order_" json:"order"`
	Parent      int    `orm:"Parent" json:"parent"`
}

func (m *Meta) TableName() string {
	return "gblog_metas"
}
