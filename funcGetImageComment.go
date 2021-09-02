package iotmakerdockerbuilder

func (e *ContainerBuilder) GetImageComment() (comment string) {

	if e.imageInspected == false {
		_, _ = e.ImageInspect()
	}

	return e.imageComment
}
