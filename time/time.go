package time

import (
	"strconv"
	"time"
)

const (
	HourTime   = "hour"
	MinuteTime = "minute"
	SecondTime = "second"
)

const (
	Year            = "2006"
	Month           = "01"
	Day             = "02"
	Hour            = "15"
	Minute          = "04"
	Second          = "05"
	Complete        = "2006-01-02 15:05:05"
	StringToTimeOne = "2006-01-02 15:04:05"
	StringToTimeTow = "2006-01-02"
)

// TodayBeginTimestamp 返回今天零点的时间戳
func TodayBeginTimestamp() int64 {
	now := time.Now()
	return now.Unix() - (int64)(now.Hour()*60*60+now.Minute()*60+now.Second())
}

// QueryBeforeTime 获取当前时间之前的某个时间点
func QueryBeforeTime(num int, flag string) string {
	currentTime := time.Now()
	deal := "-"
	var m time.Duration
	if flag == HourTime {
		m, _ = time.ParseDuration(deal + strconv.Itoa(num) + "m")
	} else if flag == MinuteTime {
		m, _ = time.ParseDuration(deal + strconv.Itoa(num) + "h")
	} else if flag == SecondTime {
		m, _ = time.ParseDuration(deal + strconv.Itoa(num) + "s")
	} else {
		return ""
	}
	result := currentTime.Add(m)
	return result.Format("2006-01-02 15:04:05")
}

// QueryAfterTime 获取当前时间之后的某个时间点
func QueryAfterTime(num int, flag string) string {
	currentTime := time.Now()
	var m time.Duration
	if flag == HourTime {
		m, _ = time.ParseDuration(strconv.Itoa(num) + "m")
	} else if flag == MinuteTime {
		m, _ = time.ParseDuration(strconv.Itoa(num) + "h")
	} else if flag == SecondTime {
		m, _ = time.ParseDuration(strconv.Itoa(num) + "s")
	} else {
		return ""
	}
	result := currentTime.Add(m)
	return result.Format("2006-01-02 15:04:05")
}

// QueryNowTime 获取当前时间
func QueryNowTime() string {
	dateTime := time.Now().Format("2006-01-02 15:04:05")
	return dateTime
}

// JudgeTimeOrder 判断一个时间是否在一个时间之后
func JudgeTimeOrder(t string) bool {
	stringTime, _ := time.Parse("2006-01-02 15:04:05", t)
	beforeOrAfter := stringTime.After(time.Now())
	return beforeOrAfter
}

// GetWeekDay 获取周几方法
func GetWeekDay(time time.Time) int {
	return int(time.Weekday())
}

// MinuteAddOrSub 时间分钟加减计算
func MinuteAddOrSub(t time.Time, num int64) time.Time {
	s := strconv.FormatInt(num, 10)
	var m time.Duration
	m, _ = time.ParseDuration(s + "m")
	return t.Add(m)
}

// HourAddOrSub 时间小时加减计算
func HourAddOrSub(t time.Time, num int64) time.Time {
	s := strconv.FormatInt(num, 10)
	var m time.Duration
	m, _ = time.ParseDuration(s + "h")
	return t.Add(m)
}

// DayAddOrSub 时间天加减计算
func DayAddOrSub(t time.Time, num int64) time.Time {
	num = num * 24
	s := strconv.FormatInt(num, 10)
	var m time.Duration
	m, _ = time.ParseDuration(s + "h")
	return t.Add(m)
}
