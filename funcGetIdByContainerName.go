package iotmakerdockerbuilder

// GetIdByContainerName
//
// English: Returns the container ID defined in SetContainerName()
//
// Português: Retorna o ID do container definido em SetContainerName()
func (e *ContainerBuilder) GetIdByContainerName() (err error) {
	e.containerID, err = e.dockerSys.ContainerFindIdByName(e.containerName)
	return
}
