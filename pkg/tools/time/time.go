package time

import (
	"strings"
	"time"
)

var (
	RFC_FORMAT  = time.RFC3339
	DATE_FORMAT = "2006-01-02 15:04:05"
)

func CompareTime(time1 string, time2 string) bool {
	t1, err := time.Parse(DATE_FORMAT, time1)
	if err != nil {
		return false
	}
	t2, err := time.Parse(DATE_FORMAT, time2)
	if err != nil {
		return false
	}
	if t1.Before(t2) {
		//处理逻辑
		return true
	}
	return false
}

func DateRFCChange(format, dateFormat, inDate string) string {
	t1, err := time.Parse(
		format,
		inDate)
	if err != nil {
		//log.Error("DateRFCChange time parse error", err.Error())
		return inDate
	}
	return t1.Format(dateFormat)
}

func GetCurrentTime() string {
	startTimestamp := time.Now().Unix()                     //获得时间戳
	return time.Unix(startTimestamp, 0).Format(DATE_FORMAT) //把时间戳转换成时间,并格式化为年月日
}

// string to time
func String2Time(timeStr string) time.Time {
	t, _ := time.ParseInLocation(DATE_FORMAT, timeStr, time.Local)
	return t
}

// string to timstamp(int64)
func String2Timestamp(timeStr string) int64 {
	if timeStr == "" {
		return 0
	}

	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(DATE_FORMAT, timeStr, loc)
	timeStamp := theTime.Unix()
	return timeStamp
}

// time to string
func Time2String(time time.Time) string {
	return time.Format(DATE_FORMAT)
}

// time to timestamp(int64)
func Time2Timestamp(time time.Time) int64 {
	return time.Unix()
}

//timestamp to time
func Timestamp2Time(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

//timestamp to string
func Timestamp2String(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(DATE_FORMAT) //把时间戳转换成时间,并格式化为年月日

}

func GormTimeFormat(timeStr string) string {
	if len(timeStr) < 20 {
		return timeStr
	}
	//return timeStr[:10] + " " + timeStr[11:19]
	return Join(timeStr[:10], " ", timeStr[11:19])
}

func Join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
