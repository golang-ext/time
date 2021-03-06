package time

import (
	"github.com/golang-ext/time/core"
	_ "unsafe"
)

//go:linkname parse core.time_parse
func parse(s, layout string, local bool) (core.Time, error)

/**
 * TODO: string => Time
 * @param s Time的字符串形式
 * @param layout 时间格式字符串
 * @param local 是否转为 Local 时间
 * @return Time
 */
func Parse(s, layout string, local bool) (core.Time, error) {
	return parse(s, layout, local)
}
