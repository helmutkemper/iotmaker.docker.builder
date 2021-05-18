package iotmakerdockerbuilder

// GetContainerID (english):
//
// GetContainerID (português): Retorna o ID do container criado
func (e *ContainerBuilder) GetContainerID() (ID string) {
	return e.containerID
}
