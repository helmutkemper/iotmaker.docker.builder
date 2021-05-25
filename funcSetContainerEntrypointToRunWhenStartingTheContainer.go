package iotmakerdockerbuilder

// SetContainerEntrypointToRunWhenStartingTheContainer (english): entrypoint to run when stating the container
//
// SetContainerEntrypointToRunWhenStartingTheContainer (português):entrypoint a ser executado quando o container inicia
func (e *ContainerBuilder) SetContainerEntrypointToRunWhenStartingTheContainer(values []string) {
	e.containerConfig.Entrypoint = values
}
