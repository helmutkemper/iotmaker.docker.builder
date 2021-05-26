package iotmakerdockerbuilder

// SetContainerAttachStandardStreamsToTty
//
// English: attach standard streams to tty
//
// Português: anexa a saída padrão do tty
func (e *ContainerBuilder) SetContainerAttachStandardStreamsToTty(value bool) {
	e.containerConfig.Tty = value
}
