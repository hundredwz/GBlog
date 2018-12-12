package config

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var (
	//Blog Config
	//base
	BlogName     string
	BlogDesc     string
	BlogKeywords string
	//comment
	BlogCommentTimeFormat = "2006-01-02 15:04"
	BlogCommentListNum    = 10
	BlogCommentAvatarUrl  = "https://cdn.v2ex.com/gravatar/"
	//read
	BlogArticleTimeFormat  = "2006-01-02"
	BlogArticleNumEachPage = 10
	BlogArticleSub         = true

	//Database Config
	DBType string
	DBName string
	DBUser string
	DBPwd  string

	//Underlying Config
	Addr      string
	Installed bool

	Config map[string]string

	Key struct {
		Pk         string
		Sk         string
		ExpireTime time.Time
		Set        bool
	}
)

func SetConfig(key, value string) {
	if Config == nil {
		Config = make(map[string]string)
	}
	Config[key] = value
}

func GetConfig(key string) string {
	if Config == nil {
		return ""
	}
	if v, ok := Config[key]; ok {
		return v
	}
	return ""
}

func InitConfig() {
	confFile := "web/admin/conf/web.conf"
	f, err := os.Open(confFile)
	if err != nil {
		return
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	for {
		line, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}
		if strings.HasPrefix(string(line), "//") {
			continue
		}
		kv := strings.Split(string(line), "=")
		if len(kv) != 2 {
			continue
		}
		name := strings.Trim(kv[0], " ")
		value := strings.Trim(kv[1], " ")
		if strings.Contains(value, "\"") {
			value = strings.Trim(value, "\"")
		}
		switch name {
		case "BlogName":
			BlogName = value
		case "BlogDesc":
			BlogDesc = value
		case "BlogKeywords":
			BlogKeywords = value
		case "BlogCommentTimeFormat":
			BlogCommentTimeFormat = value
		case "BlogCommentListNum":
			if v, err := strconv.Atoi(value); err == nil {
				BlogCommentListNum = v
			}
		case "BlogCommentAvatarUrl":
			BlogCommentAvatarUrl = value
		case "BlogArticleTimeFormat":
			BlogArticleTimeFormat = value
		case "BlogArticleNumEachPage":
			if v, err := strconv.Atoi(value); err == nil {
				BlogArticleNumEachPage = v
			}
		case "BlogArticleSub":
			if v, err := strconv.ParseBool(value); err == nil {
				BlogArticleSub = v
			}
		case "DBType":
			DBType = value
		case "DBName":
			DBName = value
		case "DBUser":
			DBUser = value
		case "DBPwd":
			DBPwd = value
		case "Addr":
			Addr = value
		case "Installed":
			if v, err := strconv.ParseBool(value); err == nil {
				Installed = v
			}
		default:
			SetConfig(name, value)
		}
	}
	return
}

func configToMap() map[string]interface{} {
	return map[string]interface{}{
		"BlogName":               BlogName,
		"BlogDesc":               BlogDesc,
		"BlogKeywords":           BlogKeywords,
		"BlogCommentTimeFormat":  BlogCommentTimeFormat,
		"BlogCommentListNum":     BlogCommentListNum,
		"BlogCommentAvatarUrl":   BlogCommentAvatarUrl,
		"BlogArticleTimeFormat":  BlogArticleTimeFormat,
		"BlogArticleNumEachPage": BlogArticleNumEachPage,
		"BlogArticleSub":         BlogArticleSub,
		"DBType":                 DBType,
		"DBName":                 DBName,
		"DBUser":                 DBUser,
		"DBPwd":                  DBPwd,
		"Addr":                   Addr,
		"Installed":              Installed,
	}
}

func UpdateConfig() error {
	t, err := template.ParseFiles("web/admin/conf/web.tpl")
	if err != nil {
		return err
	}
	webConfig := "web/admin/conf/web.conf"
	var f *os.File
	if _, err := os.Stat(webConfig); err == nil {
		f, err = os.OpenFile(webConfig, os.O_WRONLY, 0666)
		defer f.Close()
	} else {
		f, err = os.Create(webConfig)
		if err != nil {
			return err
		}
		defer f.Close()
	}
	err = t.Execute(f, configToMap())
	if err != nil {
		return err
	}
	return nil

}
