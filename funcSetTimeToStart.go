package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) SetTimeToStart(min, max time.Duration) {
	e.chaos.minimumTimeToStart = min
	e.chaos.maximumTimeToStart = max
}
