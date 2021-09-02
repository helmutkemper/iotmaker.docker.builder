package iotmakerdockerbuilder

func (e *ContainerBuilder) GetImageParent() (parent string) {

	if e.imageInspected == false {
		_, _ = e.ImageInspect()
	}

	return e.imageParent
}
