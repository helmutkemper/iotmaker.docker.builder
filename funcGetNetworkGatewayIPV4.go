package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

// GetNetworkGatewayIPV4
//
// English: Returns the gateway from the network to the IPV4 network
//
// Português: Retorna o gateway da rede para rede IPV4
func (e *ContainerBuilder) GetNetworkGatewayIPV4() (IPV4 string) {
	var err error
	var inspect iotmakerdocker.ContainerInspect

	inspect, err = e.ContainerInspect()
	if err != nil {
		util.TraceToLog()
		return
	}

	IPV4 = inspect.Network.Gateway
	return
}
