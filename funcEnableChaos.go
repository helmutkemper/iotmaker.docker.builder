package iotmakerdockerbuilder

func (e *ContainerBuilder) EnableChaos(enable bool) {
	e.chaos.enableChaos = enable
}
