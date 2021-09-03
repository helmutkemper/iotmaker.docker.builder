package iotmakerdockerbuilder

func (e *ContainerBuilder) getProbalityNumber() (probality float64) {
	return 1.0 - e.getRandSeed().Float64()
}
