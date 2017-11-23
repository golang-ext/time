package timestamp

import (
	_ "unsafe"
	"github.com/golang-ext/time.tmp/core"
)

//go:linkname parse core.timestamp_to_time
func parse(s string, _type core.TimestampType, layout string, local bool) (core.Timestamp, error)

/**
 * TODO: string => Timestamp
 * @param s Time的字符串形式
 * @param layout 时间格式字符串
 * @param local 是否转为 Local 时间
 * @return Timestamp
 */
func Parse(s string, _type core.TimestampType, layout string, local bool) (core.Timestamp, error) {
	return parse(s, _type, layout, local)
}