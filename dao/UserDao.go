package dao

import (
	"errors"
	"fmt"
	"github.com/hundredwz/GBlog/model"
)

func (db *DataBase) CreateUserTable() error {
	db.createTable(drop_gblog_users)
	if err := db.createTable(gblog_users); err != nil {
		return err
	}
	return nil
}

func (db *DataBase) AddUser(user model.User) (int, error) {
	sql := "INSERT INTO gblog_users"
	insert, args := genInsert(user)
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

func (db *DataBase) UpdateUserLogin(user model.User) error {
	update := "UPDATE gblog_users set Logged=? WHERE Uid=?"
	_, err := db.modify(update, user.Logged, user.Uid)
	return err
}

func (db *DataBase) UpdateUser(user model.User) error {
	sql := "UPDATE gblog_users "
	update, args := genUpdate(user)
	if update != "" {
		sql = sql + update + " WHERE Uid=?"
		args = append(args, user.Uid)
	}
	fmt.Println(sql, args)
	_, err := db.modify(sql, args...)
	return err
}

func (db *DataBase) DelUser(user model.User) (int, error) {
	del := "DELETE FROM gblog_users WHERE Uid=?"
	return db.modify(del, user.Uid)
}

func (db *DataBase) GetUserById(id int) model.User {
	query := "SELECT * FROM gblog_users WHERE Uid=?"
	user := model.User{}
	results := db.get(query, &model.User{}, id)
	if len(results) == 0 {
		return user
	}
	user = results[0].(model.User)
	return user
}

func (db *DataBase) UserLogin(user model.User) bool {
	query := "SELECT * FROM gblog_users WHERE Name=? AND Password=?"
	results := db.get(query, &model.User{}, user.Name, user.Password)
	if len(results) == 0 {
		return false
	}
	return true
}

func (db *DataBase) GetUserByName(Name string) model.User {
	query := "SELECT * FROM gblog_users WHERE Name=?"
	user := model.User{}
	results := db.get(query, &model.User{}, Name)
	if len(results) == 0 {
		return user
	}
	user = results[0].(model.User)
	return user
}
