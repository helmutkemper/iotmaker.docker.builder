package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) SetTimeToStartChaos(min, max time.Duration) {
	e.chaos.minimumTimeToStartChaos = min
	e.chaos.maximumTimeToStartChaos = max
}
