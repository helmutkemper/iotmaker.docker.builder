package iotmakerdockerbuilder

// GetContainerID
//
// English: Returns the ID of the created container
//
// Português: Retorna o ID do container criado
func (e *ContainerBuilder) GetContainerID() (ID string) {
	return e.containerID
}
