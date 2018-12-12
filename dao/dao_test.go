package dao

import (
	"fmt"
	"github.com/hundredwz/GBlog/model"
	"testing"
	"time"
)

func initDB() DataBase {
	db := DataBase{}
	db.Init()
	return db
}

func TestGenUpdate(t *testing.T) {
	article := model.Content{Cid: 1, Created: time.Now()}
	fmt.Println(genUpdate(article))
}
func TestGenInsert(t *testing.T) {
	article := model.Content{Cid: 1, Created: time.Now()}
	fmt.Println(genInsert(article))
}

func TestDataBase_GetComments(t *testing.T) {
	db := initDB()
	comments, err := db.GetComments(model.NewPage(1), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(comments)
}
