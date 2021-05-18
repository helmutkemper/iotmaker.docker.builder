package iotmakerdockerbuilder

// GetIdByContainerName (english):
//
// GetIdByContainerName (português): Retorna o ID do container definido em SetContainerName()
func (e *ContainerBuilder) GetIdByContainerName() (err error) {
	e.containerID, err = e.dockerSys.ContainerFindIdByName(e.containerName)
	return
}
