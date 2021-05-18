package iotmakerdockerbuilder

func (e *ContainerBuilder) ImageListExposedVolumes() (list []string, err error) {

	list, err = e.dockerSys.ImageListExposedVolumes(e.imageID)
	return
}
