package time

import (
	_ "unsafe"
	"github.com/golang-ext/time.tmp/core"
)

// TODO: 原生结构重组
//type Location core.Location
//type Time core.Time
//type Timer core.Timer
//type Duration core.Duration
//type Ticker core.Ticker
//type Month core.Month
//type Weekday core.Weekday

//go:linkname sleep core.sleep
func sleep(d core.Duration)

//go:linkname unix core.unix
func unix(sec int64, nsec int64) core.Time

//go:linkname until core.until
func until(t core.Time) core.Duration

//go:linkname new_timer core.new_timer
func new_timer(d core.Duration) *core.Timer

//go:linkname after_func core.after_func
func after_func(d core.Duration, f func()) *core.Timer

//go:linkname after core.after
func after(d core.Duration) <-chan core.Time

//go:linkname new_ticker core.new_ticker
func new_ticker(d core.Duration) *core.Ticker

//go:linkname tick core.tick
func tick(d core.Duration) <-chan core.Time

//go:linkname now core.now
func now() core.Time

//go:linkname date core.date
func date(year int, month core.Month, day, hour, min, sec, nsec int, loc *core.Location) core.Time

// TODO: 原生方法重写
func Sleep(d core.Duration) {
	sleep(d)
}
func Unix(sec int64, nsec int64) core.Time {
	return unix(sec, nsec)
}
func Until(t core.Time) core.Duration {
	return until(t)
}
func NewTimer(d core.Duration) *core.Timer {
	return new_timer(d)
}
func AfterFunc(d core.Duration, f func()) *core.Timer {
	return after_func(d, f)
}
func After(d core.Duration) <-chan core.Time {
	return after(d)
}
func NewTicker(d core.Duration) *core.Ticker {
	return new_ticker(d)
}
func Tick(d core.Duration) <-chan core.Time {
	return tick(d)
}
func Now() core.Time {
	return now()
}
func Date(year int, month core.Month, day, hour, min, sec, nsec int, loc *core.Location) core.Time {
	return date(year, month, day, hour, min, sec, nsec, loc)
}
