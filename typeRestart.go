package iotmakerdockerbuilder

import "time"

type Restart struct {
	FilterToStart      []LogFilter
	TimeToStart        Timers
	RestartProbability float64
	RestartLimit       int

	minimumEventTime time.Time
}
