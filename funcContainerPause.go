package iotmakerdockerbuilder

// ContainerPause
//
// English: pause the container
//
// Português: pausa o container
func (e *ContainerBuilder) ContainerPause() (err error) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			return
		}
	}

	err = e.dockerSys.ContainerPause(e.containerID)
	return
}
