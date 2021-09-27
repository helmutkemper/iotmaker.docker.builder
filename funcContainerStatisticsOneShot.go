package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
)

func (e *ContainerBuilder) ContainerStatisticsOneShot() (
	statsRet types.Stats,
	err error,
) {

	_, err = e.dockerSys.ContainerFindIdByName(e.containerName)
	if err != nil {
		return
	}

	statsRet, err = e.dockerSys.ContainerStatisticsOneShot(e.containerID)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
