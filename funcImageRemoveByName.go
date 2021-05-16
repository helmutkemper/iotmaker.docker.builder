package iotmaker_docker_builder

func (e *ContainerBuilder) ImageRemoveByName(name string) (err error) {
	err = e.dockerSys.ImageRemoveByName(name, false, false)
	if err != nil {
		return
	}

	return
}
