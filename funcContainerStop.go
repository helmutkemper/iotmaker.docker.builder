package iotmakerdockerbuilder

// ContainerStop
//
// English: stop the container
//
// Português: parar o container
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
