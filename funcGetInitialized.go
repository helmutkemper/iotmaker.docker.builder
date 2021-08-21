package iotmakerdockerbuilder

func (e *ContainerBuilder) GetInitialized() (initialized bool) {
	return e.init
}
