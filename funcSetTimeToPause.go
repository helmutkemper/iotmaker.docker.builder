package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) SetTimeToPause(min, max time.Duration) {
	e.chaos.minimumTimeToPause = min
	e.chaos.maximumTimeToPause = max
}
