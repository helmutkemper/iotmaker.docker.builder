package iotmakerdockerbuilder

// SetContainerEntrypointToRunWhenStartingTheContainer (english):
//
// SetContainerEntrypointToRunWhenStartingTheContainer (português):
func (e *ContainerBuilder) SetContainerEntrypointToRunWhenStartingTheContainer(values []string) {
	e.containerConfig.Entrypoint = values
}
