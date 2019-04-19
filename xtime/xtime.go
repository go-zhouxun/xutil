package xtime

import (
	"time"
)

const (
	DEFAULT_DATE      = "2006-01-02"
	DEFAULT_TIME      = "15:04:05"
	DEFAULT_DATE_TIME = "2006-01-02 15:04:05"
)

type XTime time.Time

func Now() int64 {
	return time.Now().UnixNano() / 1e6
}

func TodayDateStr() string {
	return time.Now().Format(DEFAULT_DATE)
}

func TodayTimeStr() string {
	return time.Now().Format(DEFAULT_TIME)
}

func TodayDateTimeStr() string {
	return time.Now().Format(DEFAULT_DATE_TIME)
}

func TodayStart() int64 {
	date := time.Now()
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	return date.UnixNano() / 1e6
}

func Time(days int) XTime {
	date := time.Now()
	return XTime(time.Date(date.Year(), date.Month(), date.Day()+days, 0, 0, 0, 0, date.Location()))
}

func (xtime XTime) DayStart() int64 {
	t := time.Time(xtime)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UnixNano() / 1e6
}

func (t XTime) DateStr() string {
	return time.Time(t).Format(DEFAULT_DATE)
}

func (t XTime) TimeStr() string {
	return time.Time(t).Format(DEFAULT_TIME)
}

func (t XTime) DateTimeStr() string {
	return time.Time(t).Format(DEFAULT_DATE_TIME)
}
