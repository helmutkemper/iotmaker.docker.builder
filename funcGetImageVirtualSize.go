package iotmakerdockerbuilder

func (e *ContainerBuilder) GetImageVirtualSize() (virtualSize int64) {

	if e.imageInspected == false {
		_, _ = e.ImageInspect()
	}

	return e.imageVirtualSize
}
