package util

import (
	"time"
)

func StrToDate(date string) time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return t
}
