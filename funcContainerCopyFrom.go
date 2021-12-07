package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"io"
)

func (e *ContainerBuilder) ContainerCopyFrom(
	sourcePath string,
) (
	reader io.ReadCloser,
	stats types.ContainerPathStat,
	err error,
) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	reader, stats, err = e.dockerSys.ContainerCopyFrom(e.containerID, sourcePath)
	return
}
