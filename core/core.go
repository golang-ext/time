package core

import (
	"time"
	_ "unsafe"
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
	return time.Time(t).String()
}
func (t Time) Add(d Duration) Time {
	return Time(time.Time(t).Add(time.Duration(d)))
}
func (t Time) Sub(u Time) Duration {
	return Duration(time.Time(t).Sub(time.Time(u)))
}
func (t Time) Date() (int, Month, int) {
	y, m, d := time.Time(t).Date()
	return y, Month(m), d
}
func (t Time) Day() int {
	return time.Time(t).Day()
}
func (t Time) Month() Month {
	return Month(time.Time(t).Month())
}
func (t Time) Year() int {
	return time.Time(t).Year()
}
func (t Time) Hour() int {
	return time.Time(t).Hour()
}
func (t Time) Minute() int {
	return time.Time(t).Minute()
}
func (t Time) Second() int {
	return time.Time(t).Second()
}
func (t Time) YearDay() int {
	return time.Time(t).YearDay()
}
func (t Time) Location() *Location {
	return (*Location)(time.Time(t).Location())
}
func (t Time) Local() Time {
	return Time(time.Time(t).Local())
}
func (t Time) UTC() Time {
	return Time(time.Time(t).UTC())
}
func (t Time) Weekday() Weekday {
	return Weekday(time.Time(t).Weekday())
}
func (t Time) AddDate(years, months, days int) Time {
	return Time(time.Time(t).AddDate(years, months, days))
}
func (t Time) AppendFormat(b []byte, layout string) []byte {
	return time.Time(t).AppendFormat(b, layout)
}
func (t Time) After(u Time) bool {
	return time.Time(t).After(time.Time(u))
}
func (t Time) Before(u Time) bool {
	return time.Time(t).Before(time.Time(u))
}
func (t Time) Clock() (int, int, int) {
	return time.Time(t).Clock()
}
func (t Time) GobDecode(data []byte) error {
	return time.Time(t).GobDecode(data)
}
func (t Time) GobEncode() ([]byte, error) {
	return time.Time(t).GobEncode()
}
func (t Time) Equal(u Time) bool {
	return time.Time(t).Equal(time.Time(u))
}
func (t Time) In(loc *Location) Time {
	return Time(time.Time(t).In((*time.Location)(loc)))
}
func (t Time) IsZero() bool {
	return time.Time(t).IsZero()
}
func (t Time) ISOWeek() (int, int) {
	return time.Time(t).ISOWeek()
}
func (t Time) MarshalBinary() ([]byte, error) {
	return time.Time(t).MarshalBinary()
}
func (t Time) MarshalJSON() ([]byte, error) {
	return time.Time(t).MarshalJSON()
}
func (t Time) MarshalText() ([]byte, error) {
	return time.Time(t).MarshalText()
}
func (t Time) UnmarshalBinary(data []byte) error {
	return time.Time(t).UnmarshalBinary(data)
}
func (t Time) UnmarshalJSON(data []byte) error {
	return time.Time(t).UnmarshalJSON(data)
}
func (t Time) UnmarshalText(data []byte) error {
	return time.Time(t).UnmarshalText(data)
}
func (t Time) Round(d Duration) Time {
	return Time(time.Time(t).Round(time.Duration(d)))
}
func (t Time) Zone() (string, int) {
	return time.Time(t).Zone()
}
func (t Time) Nanosecond() int {
	return time.Time(t).Nanosecond()
}
func (t Time) Truncate(d Duration) Time {
	return Time(time.Time(t).Truncate(time.Duration(d)))
}

func (l *Location) String() string {
	return (*time.Location)(l).String()
}

func (t Timer) Reset(d Duration) bool {
	return time.Timer(t).Reset(time.Duration(d))
}
func (t Timer) Stop() bool {
	return time.Timer(t).Stop()
}

func (d Duration) String() string {
	return time.Duration(d).String()
}
