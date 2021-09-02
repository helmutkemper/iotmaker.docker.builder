package iotmakerdockerbuilder

func (e *ContainerBuilder) GetImageOs() (os string) {

	if e.imageInspected == false {
		_, _ = e.ImageInspect()
	}

	return e.imageOs
}
