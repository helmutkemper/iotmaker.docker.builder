package iotmaker_docker_builder

import (
	"errors"
	"github.com/docker/docker/api/types"
	"strings"
)

func (e *ContainerBuilder) findCurrentIpAddressSupport(networkID string) (IP string, err error) {
	var res types.NetworkResource
	res, err = e.dockerSys.NetworkInspect(networkID)
	if err != nil {
		panic(err)
	}

	var pass = false
	for containerID, networkData := range res.Containers {
		if containerID == e.containerID && networkData.Name == e.containerName {
			pass = true
			var parts = strings.Split(networkData.IPv4Address, "/")
			IP = parts[0]
			return
		}
	}

	if pass == false {
		err = errors.New("container not found on bridge network")
		return
	}

	return
}
