package iotmaker_docker_builder

// SetContainerName (english):
//
// SetContainerName (português): Define o nome do container
//   value: nome do container
func (e *ContainerBuilder) SetContainerName(value string) {
	e.containerName = value
}
