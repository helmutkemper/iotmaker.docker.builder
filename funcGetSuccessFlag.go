package iotmakerdockerbuilder

func (e *ContainerBuilder) GetSuccessFlag() (success bool) {
	return e.chaos.foundSuccess
}
