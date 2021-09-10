package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) selectBetweenMaxAndMin(max, min time.Duration) (selected time.Duration) {
	randValue := e.getRandSeed().Int63n(int64(max)-int64(min)) + int64(min)
	return time.Duration(randValue)
}
