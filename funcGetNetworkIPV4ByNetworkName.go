package iotmaker_docker_builder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

func (e *ContainerBuilder) GetNetworkIPV4ByNetworkName(networkName string) (IPV4 string) {
	var err error
	var inspect iotmakerdocker.ContainerInspect

	if e.network == nil {
		networkName = "bridge"
	} else {
		networkName = e.network.GetNetworkName()
	}

	inspect, err = e.ContainerInspect()
	if err != nil {
		return
	}

	IPV4 = inspect.Network.Networks[networkName].IPAddress
	return
}
