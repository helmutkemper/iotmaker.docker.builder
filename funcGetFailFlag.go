package iotmakerdockerbuilder

func (e *ContainerBuilder) GetFailFlag() (fail bool) {
	return e.chaos.foundFail
}
