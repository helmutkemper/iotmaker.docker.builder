package iotmakerdockerbuilder

func (e *ContainerBuilder) GetImageVariant() (variant string) {

	if e.imageInspected == false {
		_, _ = e.ImageInspect()
	}

	return e.imageVariant
}
