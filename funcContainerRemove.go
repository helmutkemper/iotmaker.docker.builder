package iotmakerdockerbuilder

// ContainerRemove (english):
//
// ContainerRemove (português): parar e remover o container
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
