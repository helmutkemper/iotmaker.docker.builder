package iotmakerdockerbuilder

func (e *ContainerBuilder) SetRestartProbability(restartProbability, restartChangeIpProbability float64, limit int) {
	e.chaos.restartProbability = restartProbability
	e.chaos.restartChangeIpProbability = restartChangeIpProbability
	e.chaos.restartLimit = limit
}
