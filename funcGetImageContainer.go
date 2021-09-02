package iotmakerdockerbuilder

func (e *ContainerBuilder) GetImageContainer() (container string) {

	if e.imageInspected == false {
		_, _ = e.ImageInspect()
	}

	return e.imageContainer
}
