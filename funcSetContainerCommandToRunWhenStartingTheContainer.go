package iotmakerdockerbuilder

// SetContainerCommandToRunWhenStartingTheContainer
//
// English: command to run when stating the container (style Dockerfile CMD)
//
// Português: comando a ser executado quando o container inicia (estilo Dockerfile CMD)
func (e *ContainerBuilder) SetContainerCommandToRunWhenStartingTheContainer(values []string) {
	e.containerConfig.Cmd = values
}
