package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) SetTimeToRestart(min, max time.Duration) {
	e.chaos.minimumTimeToRestart = min
	e.chaos.maximumTimeToRestart = max
}
