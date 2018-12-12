package model

import "time"

type User struct {
	Uid        int       `orm:"Uid"`
	Name       string    `orm:"Name"`
	Password   string    `orm:"Password"`
	Mail       string    `orm:"Mail"`
	Url        string    `orm:"Url"`
	ScreenName string    `orm:"ScreenName"`
	Created    time.Time `orm:"Created"`
	Logged     time.Time `orm:"Logged"`
	Permission string    `orm:"Permission"`
}

func (u *User) TableName() string {
	return "gblog_users"
}
