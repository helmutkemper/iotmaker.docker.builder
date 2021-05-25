package iotmakerdockerbuilder

// GetChannelOnContainerInspect (english):
//
// GetChannelOnContainerInspect (português): Canas disparado a cada ciclo do ticker definido em SetInspectInterval()
func (e *ContainerBuilder) GetChannelOnContainerInspect() (channel *chan bool) {
	return e.onContainerInspect
}
