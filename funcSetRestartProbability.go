package iotmakerdockerbuilder

func (e *ContainerBuilder) SetRestartProbability(probability float64, limit int) {
	e.chaos.restartProbability = probability
	e.chaos.restartLimit = limit
}
