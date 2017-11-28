package core

import (
	_ "unsafe"
)

type TimestampType uint
type Timestamp int64

const (
	T_NanoSecond  TimestampType = iota
	T_MicroSecond
	T_MilliSecond
	T_Second
)

//go:linkname timestamp_parse core.timestamp_to_time
func timestamp_parse(s string, _type TimestampType, layout string, local bool) (Timestamp, error) {
	_time, err := time_parse(s, layout, local)
	switch _type {
	case T_Second:
		return Timestamp(_time.Unix()), err
	case T_MilliSecond:
		return Timestamp(_time.UnixMilli()), err
	case T_MicroSecond:
		return Timestamp(_time.UnixMicro()), err
	case T_NanoSecond:
		return Timestamp(_time.UnixNano()), err
	default:
		return Timestamp(_time.Unix()), err
	}
}

/**
 * TODO: 时间戳 => Time
 * @param _type 时间戳类型
 * @return 时间
 */
func (timestamp Timestamp) Time(_type TimestampType) Time {
	switch _type {
	case T_Second:
		return unix(int64(timestamp), 0)
	case T_MilliSecond:
		return unix(0, int64(timestamp*1000*1000))
	case T_MicroSecond:
		return unix(0, int64(timestamp*1000))
	case T_NanoSecond:
		return unix(0, int64(timestamp))
	default:
		return unix(int64(timestamp), 0)
	}
}

/**
 * TODO: 时间戳 => string
 * @param layout 时间格式字符串
 * @param _type 时间戳类型
 * @return string
 */
func (timestamp Timestamp) Format(layout string, _type TimestampType) string {
	switch _type {
	case T_Second:
		return unix(int64(timestamp), 0).Format(layout)
	case T_MilliSecond:
		return unix(0, int64(timestamp*1000*1000)).Format(layout)
	case T_MicroSecond:
		return unix(0, int64(timestamp*1000)).Format(layout)
	case T_NanoSecond:
		return unix(0, int64(timestamp)).Format(layout)
	default:
		return unix(int64(timestamp), 0).Format(layout)
	}
}

/**
 * TODO: 时间戳相加
 * @param timestamp
 * @return int64
 */
func (timestamp Timestamp) Add(_t Timestamp) int64 {
	return int64(timestamp + _t)
}

/**
 * TODO: 时间戳相减
 * @param timestamp
 * @return int64
 */
func (timestamp Timestamp) Sub(_t Timestamp) int64 {
	return int64(timestamp) - int64(_t)
}
