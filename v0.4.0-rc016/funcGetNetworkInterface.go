package iotmakerdockerbuilder

import (
	isolatedNetwork "github.com/helmutkemper/iotmaker.docker.builder.network.interface"
)

// GetNetworkInterface (english):
//
// GetNetworkInterface (português):
func (e *ContainerBuilder) GetNetworkInterface() (network isolatedNetwork.ContainerBuilderNetworkInterface) {
	return e.network
}
