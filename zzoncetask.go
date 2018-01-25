package cron

import (
	"time"
)

// OnceDelaySchedule represent schedule just run once
type OnceDelaySchedule struct {
	Start time.Time
	Delay time.Duration
}

// Once just once
func Once(duration time.Duration) OnceDelaySchedule {
	if duration < time.Second {
		duration = time.Second
	}
	now := time.Now()
	return OnceDelaySchedule{
		Delay: duration - time.Duration(duration.Nanoseconds())%time.Second,
		Start: now.Add(-time.Duration(now.Nanosecond()) * time.Nanosecond),
	}
}

// Next impl Schedule
func (once OnceDelaySchedule) Next(t time.Time) time.Time {
	if once.Start.Add(once.Delay - time.Duration(t.Nanosecond())*time.Nanosecond).Before(t) {
		return time.Time{}
	}
	return t.Add(once.Delay - time.Duration(t.Nanosecond())*time.Nanosecond)
}
