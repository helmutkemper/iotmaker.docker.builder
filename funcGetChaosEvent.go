package iotmakerdockerbuilder

func (e *ContainerBuilder) GetChaosEvent() (eventChannel chan Event) {
	return e.chaos.event
}
