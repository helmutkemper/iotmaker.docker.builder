package iotmakerdockerbuilder

// SetContainerName
//
// English: Defines the name of the container
//
//   value: container name
//
// Português: Define o nome do container
//
//   value: nome do container
func (e *ContainerBuilder) SetContainerName(value string) {
	e.containerName = value
}
