package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
)

// ContainerStatisticsOneShot
//
// English: Returns the container's memory and system consumption data at the time of the query.
//
// Português: Retorna os dados de consumo de memória e sistema do container no instante da consulta.
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
