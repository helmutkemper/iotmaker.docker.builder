package iotmaker_docker_builder

// ContainerStop (english):
//
// ContainerStop (português): parar o container
func (e *ContainerBuilder) ContainerStop() (err error) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			return
		}
	}

	err = e.dockerSys.ContainerStop(e.containerID)
	return
}
