package iotmaker_docker_builder

// ContainerRemove (english):
//
// ContainerRemove (português): para e remove o container
func (e *ContainerBuilder) ContainerRemove(removeVolumes bool) (err error) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			return
		}
	}

	err = e.dockerSys.ContainerStopAndRemove(e.containerID, removeVolumes, false, false)
	return
}
