package model

type CMR struct {
	Cid int `orm:"Cid" json:"mid"`
	Mid int `orm:"Mid" json:"cid"`
}
