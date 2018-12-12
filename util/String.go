package util

import (
	"time"
)

var (
	YMD   = "2006-01-02"
	YMDHM = "2006-01-02 15:04"
	MDHM  = "01-02 15:04"
)

func SubString(str string, size int) string {
	if len([]rune(str)) < size {
		return str
	}
	return string([]rune(str)[:size])
}

func FormatTime(t time.Time, format string) string {
	return t.Format(format)
}
