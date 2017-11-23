package core

import (
	_ "unsafe"
	"time"
)

// 重写原生 结构
type Location time.Location
type Time time.Time
type Timer time.Timer
type Duration time.Duration
type Ticker time.Ticker
type Month time.Month
type Weekday time.Weekday

var (
	Local = (*Location)(time.Local)
	UTC   = (*Location)(time.UTC)
)

// time.Duration
const (
	Nanosecond  = Duration(time.Nanosecond)
	Microsecond = Duration(time.Microsecond)
	Millisecond = Duration(time.Millisecond)
	Second      = Duration(time.Second)
	Minute      = Duration(time.Minute)
	Hour        = Duration(time.Hour)
	Day         = Hour * 24
	Week        = Day * 7
)

// time.WeekDay
const (
	Monday    = Weekday(time.Monday)
	Tuesday   = Weekday(time.Tuesday)
	Wednesday = Weekday(time.Wednesday)
	Thursday  = Weekday(time.Thursday)
	Friday    = Weekday(time.Friday)
	Saturday  = Weekday(time.Saturday)
	Sunday    = Weekday(time.Sunday)
)

// time.Month
const (
	January   = Month(time.January)
	February  = Month(time.February)
	March     = Month(time.March)
	April     = Month(time.April)
	May       = Month(time.May)
	June      = Month(time.June)
	July      = Month(time.July)
	August    = Month(time.August)
	September = Month(time.September)
	October   = Month(time.October)
	November  = Month(time.November)
	December  = Month(time.December)
)

//go:linkname sleep core.sleep
func sleep(d Duration) {
	time.Sleep(time.Duration(d))
}

//go:linkname unix core.unix
func unix(sec int64, nsec int64) Time {
	return Time(time.Unix(sec, nsec))
}

//go:linkname until core.until
func until(t Time) Duration {
	return Duration(time.Until(time.Time(t)))
}

//go:linkname new_timer core.new_timer
func new_timer(d Duration) *Timer {
	return (*Timer)(time.NewTimer(time.Duration(d)))
}

//go:linkname after_func core.after_func
func after_func(d Duration, f func()) *Timer {
	return (*Timer)(time.AfterFunc(time.Duration(d), f))
}

//go:linkname after core.after
func after(d Duration) <-chan Time {
	ch := make(chan Time)
	ch <- Time(<-time.After(time.Duration(d)))
	return ch
}

//func Parse(layout, value string) (Time, error) {
//	t, err := time.Parse(layout, value)
//	return Time{t}, err
//}
//go:linkname new_ticker core.new_ticker
func new_ticker(d Duration) *Ticker {
	return (*Ticker)(time.NewTicker(time.Duration(d)))
}

//go:linkname tick core.tick
func tick(d Duration) <-chan Time {
	ch := make(chan Time)
	ch <- Time(<-time.Tick(time.Duration(d)))
	return ch
}

//go:linkname now core.now
func now() Time {
	return Time(time.Now())
}

//go:linkname date core.date
func date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time {
	return Time(time.Date(year, time.Month(month), day, hour, min, sec, nsec, (*time.Location)(loc)))
}

func (t Time) String() string {
	return t.Format("2006-01-02 15:04:05.999999999 -0700 MST")
}
func (t Time) Add(d Duration) Time {
	return Time(time.Time(t).Add(time.Duration(d)))
}

func (l *Location) String() string {
	return (*time.Location)(l).String()
}
