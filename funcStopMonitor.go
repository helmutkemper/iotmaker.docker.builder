package iotmakerdockerbuilder

func (e *ContainerBuilder) StopMonitor() {
	if e.chaos.monitorStop == nil {
		return
	}

	if len(e.chaos.monitorStop) == 0 {
		e.chaos.monitorStop <- struct{}{}
	}
}
