package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) GetImageCreated() (created time.Time) {

	if e.imageInspected == false {
		_, _ = e.ImageInspect()
	}

	return e.imageCreated
}
