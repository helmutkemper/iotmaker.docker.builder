package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

// GetNetworkGatewayIPV4 (english):
//
// GetNetworkGatewayIPV4 (português): Retorna o gateway da rede para rede IPV4
func (e *ContainerBuilder) GetNetworkGatewayIPV4() (IPV4 string) {
	var err error
	var inspect iotmakerdocker.ContainerInspect

	inspect, err = e.ContainerInspect()
	if err != nil {
		return
	}

	IPV4 = inspect.Network.Gateway
	return
}
