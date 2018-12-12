package tests

import (
	"fmt"
	"github.com/hundredwz/GBlog/model"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestReflect(t *testing.T) {
	a := model.Content{}
	r := reflect.TypeOf(a)
	aValue := reflect.ValueOf(a)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	if r.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return
	}
	insert := "INSERT INTO A("
	value := " VALUES("
	fieldNum := r.NumField()
	for i := 0; i < fieldNum; i++ {
		fieldName := r.Field(i).Name
		//fmt.Println(r.Field(i).Type)
		switch aValue.FieldByName(fieldName).Interface().(type) {
		case int:
			if aValue.FieldByName(fieldName).Int() != -1 {
				insert = insert + fieldName + ","
				value = fmt.Sprintf("%v%v,", value, aValue.FieldByName(fieldName).Int())
			}
		case string:
			if aValue.FieldByName(fieldName).String() != "" {
				insert = insert + fieldName + ","
				value = fmt.Sprintf("%v,%v,", value, aValue.FieldByName(fieldName).Int())
			}
			fmt.Println("string")
		case time.Time:
			fmt.Println("time")
		default:
			fmt.Println("not known")
		}

		//res := aValue.FieldByName(fieldName).Interface()
		//fmt.Println(res)

	}
	fmt.Println(insert, value)
}
