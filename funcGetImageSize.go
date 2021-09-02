package iotmakerdockerbuilder

func (e *ContainerBuilder) GetImageSize() (size int64) {

	if e.imageInspected == false {
		_, _ = e.ImageInspect()
	}

	return e.imageSize
}
