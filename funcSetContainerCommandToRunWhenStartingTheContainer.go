package iotmakerdockerbuilder

// SetContainerCommandToRunWhenStartingTheContainer (english):
//
// SetContainerCommandToRunWhenStartingTheContainer (português):
func (e *ContainerBuilder) SetContainerCommandToRunWhenStartingTheContainer(values []string) {
	e.containerConfig.Cmd = values
}
