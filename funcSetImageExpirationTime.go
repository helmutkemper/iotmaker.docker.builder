package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) SetImageExpirationTime(expiration time.Duration) {
	e.imageExpirationTime = expiration
}
