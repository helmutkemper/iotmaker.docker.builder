package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) GetImageExpirationTime() (expiration time.Duration) {
	return e.imageExpirationTime
}
