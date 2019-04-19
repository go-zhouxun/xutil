package xtime

import (
	"time"
)

const (
	DEFAULT_DATE      = "2006-01-02"
	DEFAULT_TIME      = "15:04:05"
	DEFAULT_DATE_TIME = "2006-01-02 15:04:05"
)

func Now() int64 {
	return time.Now().UnixNano()/1e6
}

func DateStr() string {
	return time.Now().Format(DEFAULT_DATE)
}

func TimeStr() string {
	return time.Now().Format(DEFAULT_TIME)
}

func DateTimeStr() string {
	return time.Now().Format(DEFAULT_DATE_TIME)
}

func TodayStart() int64 {
	date := time.Now()
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	return date.UnixNano() / 1e6
}

func Date(days int) time.Time {
	date := time.Now()
	return time.Date(date.Year(), date.Month(), date.Day()+days, 0, 0, 0, 0, date.Location())
}

func DayStart(date time.Time) int64 {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	return date.UnixNano() / 1e6
}

