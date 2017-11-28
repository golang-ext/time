package time

import (
	"github.com/golang-ext/time/core"
	"time"
)

var (
	Local = core.Local
	UTC   = core.UTC
)

// time.Duration
const (
	Nanosecond  = core.Nanosecond
	Microsecond = core.Microsecond
	Millisecond = core.Millisecond
	Second      = core.Second
	Minute      = core.Minute
	Hour        = core.Hour
	Day         = core.Day
	Week        = core.Week
)

// time.WeekDay
const (
	Monday    = core.Monday
	Tuesday   = core.Tuesday
	Wednesday = core.Wednesday
	Thursday  = core.Thursday
	Friday    = core.Friday
	Saturday  = core.Saturday
	Sunday    = core.Sunday
)

// time.Month
const (
	January   = core.January
	February  = core.February
	March     = core.March
	April     = core.April
	May       = core.May
	June      = core.June
	July      = core.July
	August    = core.August
	September = core.September
	October   = core.October
	November  = core.November
	December  = core.December
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
