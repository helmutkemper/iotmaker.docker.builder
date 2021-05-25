package iotmakerdockerbuilder

// SetContainerAttachStandardStreamsToTty (english):
//
// SetContainerAttachStandardStreamsToTty (português):
func (e *ContainerBuilder) SetContainerAttachStandardStreamsToTty(value bool) {
	e.containerConfig.Tty = value
}
