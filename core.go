package time

import (
	"time"
	"github.com/golang-ext/time/core"
)

var (
	Local = core.Location(time.Local)
	UTC   = core.Location(time.UTC)
)

// time.Duration
const (
	Nanosecond  = core.Duration(time.Nanosecond)
	Microsecond = core.Duration(time.Microsecond)
	Millisecond = core.Duration(time.Millisecond)
	Second      = core.Duration(time.Second)
	Minute      = core.Duration(time.Minute)
	Hour        = core.Duration(time.Hour)
	Day         = Hour * 24
	Week        = Day * 7
)

// time.WeekDay
const (
	Monday    = core.Weekday(time.Monday)
	Tuesday   = core.Weekday(time.Tuesday)
	Wednesday = core.Weekday(time.Wednesday)
	Thursday  = core.Weekday(time.Thursday)
	Friday    = core.Weekday(time.Friday)
	Saturday  = core.Weekday(time.Saturday)
	Sunday    = core.Weekday(time.Sunday)
)

// time.Month
const (
	January   = core.Month(time.January)
	February  = core.Month(time.February)
	March     = core.Month(time.March)
	April     = core.Month(time.April)
	May       = core.Month(time.May)
	June      = core.Month(time.June)
	July      = core.Month(time.July)
	August    = core.Month(time.August)
	September = core.Month(time.September)
	October   = core.Month(time.October)
	November  = core.Month(time.November)
	December  = core.Month(time.December)
)

// time.layout
const (
	ANSIC       = time.ANSIC
	UnixDate    = time.UnixDate
	RubyDate    = time.RubyDate
	RFC822      = time.RFC822
	RFC822Z     = time.RFC1123Z
	RFC850      = time.RFC850
	RFC1123     = time.RFC1123
	RFC1123Z    = time.RFC1123Z
	RFC3339     = time.RFC3339
	RFC3339Nano = time.RFC3339Nano
	Kitchen     = time.Kitchen
	// Handy time stamps.
	Stamp      = time.Stamp
	StampMilli = time.StampMilli
	StampMicro = time.StampMicro
	StampNano  = time.StampNano

	LAYOUT_DEFALUT       = "2006-01-02 15:04:05"
	LAYOUT_MILLI_DEFAULT = "2006-01-02 15:04:05.000"
	LAYOUT_MICRO_DEFAULT = "2006-01-02 15:04:05.000000"
	LAYOUT_NANO_DEFAULT  = "2006-01-02 15:04:05.000000000"
)

const (
	年  = "2006"
	月  = "01"
	日  = "02"
	时  = "15"
	分  = "04"
	秒  = "05"
	星期 = "Mon"
	时区 = "MST"
	纳秒 = ".000000000"
	微秒 = ".000000"
	毫秒 = ".000"

	时区_  = "-0700"
	时区__ = "Z07:00"
	时_   = "3"
	午    = "PM"
	年_   = "06"
	月_   = "Jan"
	日_   = "_2"
	纳秒_  = ".999999999"
)
