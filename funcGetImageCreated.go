package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) GetImageCreated() (created time.Time) {
	return e.imageCreated
}
