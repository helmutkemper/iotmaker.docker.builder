package iotmaker_docker_builder

import (
	"errors"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

func (e *ContainerBuilder) GetNetworkIPV4ByNetworkName(networkName string) (IPV4 string, err error) {
	var found bool
	var inspect iotmakerdocker.ContainerInspect

	inspect, err = e.ContainerInspect()
	if err != nil {
		return
	}

	_, found = inspect.Network.Networks[networkName]
	if found == false {
		err = errors.New("network name not found")
		return
	}

	IPV4 = inspect.Network.Networks[networkName].IPAddress
	return
}
