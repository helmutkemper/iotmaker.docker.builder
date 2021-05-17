package iotmaker_docker_builder

func (e *ContainerBuilder) ImageListExposedVolumes() (list []string, err error) {

	list, err = e.dockerSys.ImageListExposedVolumes(e.imageID)
	return
}
