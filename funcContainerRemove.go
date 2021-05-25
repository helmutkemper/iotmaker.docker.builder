package iotmakerdockerbuilder

// ContainerRemove
//
// English: stop and remove the container
//
//   removeVolumes: removes docker volumes linked to the container
//
// Português: parar e remover o container
//
//   removeVolumes: remove os volumes docker vinculados ao container
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
