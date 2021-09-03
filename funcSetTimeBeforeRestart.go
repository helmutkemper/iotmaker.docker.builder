package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) SetTimeBeforeRestart(min, max time.Duration) {
	e.chaos.minimumTimeBeforeRestart = min
	e.chaos.maximumTimeBeforeRestart = max
}
