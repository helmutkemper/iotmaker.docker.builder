package iotmakerdockerbuilder

func (e *ContainerBuilder) GetImageAuthor() (author string) {

	if e.imageInspected == false {
		_, _ = e.ImageInspect()
	}

	return e.imageAuthor
}
