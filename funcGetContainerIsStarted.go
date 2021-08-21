package iotmakerdockerbuilder

func (e ContainerBuilder) GetContainerIsStarted() (started bool) {
	return e.startedAfterBuild
}
