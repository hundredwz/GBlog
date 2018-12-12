package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hundredwz/GBlog/config"
	"reflect"
	"strings"
)

const (
	timeFormate = "2006-01-02 15:04:05"
)

type DataBase struct {
	dataSource string
	conn       *sql.DB
}

func (db *DataBase) Init() {
	db.dataSource = fmt.Sprintf("%s:%s@/%s?charset=utf8&loc=Asia%%2FShanghai&parseTime=true", config.DBUser, config.DBPwd, config.DBName)
}

func (db *DataBase) Open() error {
	if db.conn != nil {
		return nil
	}
	var err error
	db.conn, err = sql.Open("mysql", db.dataSource)
	return err
}

func (db *DataBase) Close() {
	if db.conn != nil {
		db.conn.Close()
	}
	db.conn = nil
}

func (db *DataBase) Connection() error {
	db.Init()
	err := db.Open()
	defer db.Close()
	if err != nil {
		return err
	}
	err = db.conn.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) createTable(table string) error {
	db.Open()
	defer db.Close()
	_, err := db.conn.Exec(table)
	return err
}

func (db *DataBase) get(query string, varType interface{}, cond ...interface{}) []interface{} {
	db.Open()
	defer db.Close()
	stmt, err := db.conn.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(cond...)
	if err != nil {
		return nil
	}
	defer rows.Close()
	return getRows(rows, varType)
}

func (db *DataBase) modify(modify string, cond ...interface{}) (int, error) {
	db.Open()
	defer db.Close()
	stmt, err := db.conn.Prepare(modify)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(cond...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func getRows(rows *sql.Rows, varType interface{}) []interface{} {
	result := make([]interface{}, 0)
	s := reflect.ValueOf(varType).Elem()
	l := s.NumField()
	oneRow := make([]interface{}, l)
	for i := 0; i < l; i++ {
		oneRow[i] = s.Field(i).Addr().Interface()
	}
	for rows.Next() {
		rows.Scan(oneRow...)
		result = append(result, s.Interface())
	}
	return result
}

func genUpdate(varType interface{}) (string, []interface{}) {
	update := " SET "
	args := make([]interface{}, 0)
	v := reflect.ValueOf(varType)
	t := reflect.TypeOf(varType)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return "", args
	}
	elemSize := t.NumField()
	if elemSize == 0 {
		return "", nil
	}
	for i := 0; i < elemSize; i++ {
		fieldName := t.Field(i).Name
		tagName := t.Field(i).Tag.Get("orm")
		if tagName == "" {
			continue
		}
		arg := validUpdateType(v.FieldByName(fieldName))
		if arg == nil {
			continue
		}
		update = update + tagName + "=?,"
		args = append(args, v.FieldByName(fieldName).Interface())
	}
	update = strings.TrimSuffix(update, ",")
	return update, args
}

func genInsert(varType interface{}) (string, []interface{}) {
	insert := "("
	values := " VALUES("
	args := make([]interface{}, 0)
	v := reflect.ValueOf(varType)
	t := reflect.TypeOf(varType)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return "", args
	}
	elemSize := t.NumField()
	if elemSize == 0 {
		return "", nil
	}
	for i := 0; i < elemSize; i++ {
		fieldName := t.Field(i).Name
		tagName := t.Field(i).Tag.Get("orm")
		if tagName == "" {
			continue
		}
		arg := validUpdateType(v.FieldByName(fieldName))
		if arg == nil {
			continue
		}
		insert = insert + tagName + ","
		values = values + "?,"
		args = append(args, v.FieldByName(fieldName).Interface())
	}
	insert = strings.TrimSuffix(insert, ",") + ")"
	values = strings.TrimSuffix(values, ",") + ")"
	insert = insert + values
	return insert, args
}

func validUpdateType(v reflect.Value) interface{} {
	if !v.IsValid() {
		return nil
	}
	switch v.Kind() {
	case reflect.String:
		if v.String() == "" {
			return nil
		}
		return v.String()
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v.Uint() == 0 {
			return nil
		}
		return v.Uint()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.Int() == 0 {
			return nil
		}
		return v.Int()
	case reflect.Float32, reflect.Float64:
		if v.Float() == 0 {
			return nil
		}
		return v.Float()
	case reflect.Struct:
		switch v.Type().String() {
		case "time.Time":
			m := v.MethodByName("Format")
			rets := m.Call([]reflect.Value{reflect.ValueOf(timeFormate)})
			if strings.Contains(rets[0].String(), "0001") {
				return nil
			}
			return v
		default:
			return v
		}
	default:
		return v
	}
}
