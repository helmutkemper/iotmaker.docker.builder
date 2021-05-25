package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

// GetNetworkIPV4
//
// English: Return the IPV4 from the docker network
//
// Português: Retorno o IPV4 da rede do docker
func (e *ContainerBuilder) GetNetworkIPV4() (IPV4 string) {
	var err error
	var inspect iotmakerdocker.ContainerInspect

	inspect, err = e.ContainerInspect()
	if err != nil {
		return
	}

	IPV4 = inspect.Network.IPAddress
	return
}
