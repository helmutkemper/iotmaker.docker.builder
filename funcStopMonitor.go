package iotmakerdockerbuilder

func (e *ContainerBuilder) StopMonitor() {
	if e.chaos.monitorStop == nil {
		return
	}

	e.chaos.monitorStop <- struct{}{}
}
