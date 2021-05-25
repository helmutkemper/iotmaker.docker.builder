package iotmakerdockerbuilder

import (
	"errors"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

// GetNetworkGatewayIPV4ByNetworkName
//
// English: If the container is connected to more than one network, this function returns the gateway of the chosen
// network.
//
//   Note: the default docker network is named "bridge"
//
// Português: Caso o container esteja ligado em mais de uma rede, esta função devolve o gateway da rede escolhida.
//
//   Nota: a rede padrão do docker tem o nome "bridge"
func (e *ContainerBuilder) GetNetworkGatewayIPV4ByNetworkName(networkName string) (IPV4 string, err error) {
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

	IPV4 = inspect.Network.Networks[networkName].Gateway
	return
}
