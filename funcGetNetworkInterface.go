package iotmakerdockerbuilder

import (
	isolatedNetwork "github.com/helmutkemper/iotmaker.docker.builder.network.interface"
)

// GetNetworkInterface
//
// English: Returns the object defined for the network control
//
// Português: Retorna o objeto definido para o controle da rede
func (e *ContainerBuilder) GetNetworkInterface() (network isolatedNetwork.ContainerBuilderNetworkInterface) {
	return e.network
}
