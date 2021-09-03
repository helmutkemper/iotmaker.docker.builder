package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) SetTimeToUnpause(min, max time.Duration) {
	e.chaos.minimumTimeToUnpause = min
	e.chaos.maximumTimeToUnpause = max
}
