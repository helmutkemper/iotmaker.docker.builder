package iotmakerdockerbuilder

import (
	isolatedNetwork "github.com/helmutkemper/iotmaker.docker.builder.network.interface"
)

// SetNetworkDocker (english):
//
// SetNetworkDocker (português): Define o ponteiro do gerenciador de rede docker
//   network: ponteiro para o objeto gerenciador de rede.
//
//     Nota: compatível com o objeto dockerBuilderNetwork.ContainerBuilderNetwork{}
func (e *ContainerBuilder) SetNetworkDocker(network isolatedNetwork.ContainerBuilderNetworkInterface) {
	e.network = network
}
