package iotmakerdockerbuilder

// SetContainerEntrypointToRunWhenStartingTheContainer
//
// English: entrypoint to run when stating the container
//
// Português: entrypoint a ser executado quando o container inicia
func (e *ContainerBuilder) SetContainerEntrypointToRunWhenStartingTheContainer(values []string) {
	e.containerConfig.Entrypoint = values
}
