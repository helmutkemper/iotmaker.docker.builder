package iotmakerdockerbuilder

// SetContainerCommandToRunWhenStartingTheContainer (english): command to run when stating the container
//
// SetContainerCommandToRunWhenStartingTheContainer (português): comando a ser executado quando o container inicia
func (e *ContainerBuilder) SetContainerCommandToRunWhenStartingTheContainer(values []string) {
	e.containerConfig.Cmd = values
}
