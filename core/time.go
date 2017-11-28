package core

import (
	"time"
	_ "unsafe"
)

//go:linkname time_parse core.time_parse
func time_parse(s, layout string, local bool) (Time, error) {
	if local {
		theTime, err := time.ParseInLocation(layout, s, time.Local) //使用模板在对应时区转化为time.time类型
		return Time(theTime), err
	} else {
		theTime, err := time.Parse(layout, s) //使用模板在对应时区转化为time.time类型
		return Time(theTime), err
	}
}

/**
 * TODO: 获取 19 位 时间戳
 * @param t time.Time
 * @return int64
 */
func (t Time) UnixNano() Timestamp {
	return Timestamp(time.Time(t).UnixNano())
}

/**
 * TODO: 获取 16 位 时间戳
 * @param t time.Time
 * @return int64
 */
func (t Time) UnixMicro() Timestamp {
	return Timestamp(t.UnixNano() / 1000)
}

/**
 * TODO: 获取 13 位 时间戳
 * @param t time.Time
 * @return int64
 */
func (t Time) UnixMilli() Timestamp {
	return Timestamp(t.UnixMicro() / 1000)
}

/**
 * TODO: 获取 10 位 时间戳
 * @param t time.Time
 * @return int64
 */
func (t Time) Unix() Timestamp {
	return Timestamp(time.Time(t).Unix())
}

/**
 * TODO: Time => string
 * @param layout 时间格式字符串
 * @return string
 */
func (t Time) Format(layout string) string {
	return time.Time(t).Format(layout)
}
