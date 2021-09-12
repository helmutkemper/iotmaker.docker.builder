package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types/network"
	"github.com/helmutkemper/util"
)

func (e *ContainerBuilder) NetworkChangeIp() (err error) {
	var networkID = e.network.GetNetworkID()
	err = e.dockerSys.NetworkDisconnect(networkID, e.containerID, false)
	if err != nil {
		util.TraceToLog()
		return
	}

	var netConfig *network.NetworkingConfig
	e.IPV4Address, netConfig, err = e.network.GetConfiguration()
	if err != nil {
		util.TraceToLog()
		return
	}

	err = e.dockerSys.NetworkConnect(networkID, e.containerID, netConfig.EndpointsConfig[e.network.GetNetworkName()])
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
