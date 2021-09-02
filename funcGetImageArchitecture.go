package iotmakerdockerbuilder

func (e *ContainerBuilder) GetImageArchitecture() (architecture string) {

	if e.imageInspected == false {
		_, _ = e.ImageInspect()
	}

	return e.imageArchitecture
}
