package iotmakerdockerbuilder

func (e *ContainerBuilder) GetImageOsVersion() (osVersion string) {

	if e.imageInspected == false {
		_, _ = e.ImageInspect()
	}

	return e.imageOsVersion
}
