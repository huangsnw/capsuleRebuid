package utils

import (
	"capsuleRebuild/config"
	"time"
)

func FormatString2Time(timeStr string) time.Time {
	dateFormat := "2006-01-02"
	t, e := time.ParseInLocation(dateFormat, timeStr, time.Local)
	config.PrintError(e, "time format change faild")
	return t
}

func NextDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}
