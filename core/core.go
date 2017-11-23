package core

import (
	_ "unsafe"
	"time"
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
	return &Timer(time.NewTimer(time.Duration(d)))
}

//go:linkname after_func core.after_func
func after_func(d Duration, f func()) *Timer {
	return &Timer(time.AfterFunc(time.Duration(d), f))
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
	return &Ticker(time.NewTicker(time.Duration(d)))
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
	return Time(time.Date(year, time.Month(month), day, hour, min, sec, nsec, &time.Location(loc)))
}
