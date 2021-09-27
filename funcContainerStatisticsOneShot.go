package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
)

func (e *ContainerBuilder) ContainerStatisticsOneShot() (
	statsRet types.Stats,
	err error,
) {

	statsRet, err = e.dockerSys.ContainerStatisticsOneShot(e.containerID)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
