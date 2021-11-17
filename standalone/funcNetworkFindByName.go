package standalone

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

func NetworkFindByName(name string) (id string, err error) {
	dockerSys := iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	id, err = dockerSys.NetworkFindIdByName(name)
	return
}
