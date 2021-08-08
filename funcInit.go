package iotmakerdockerbuilder

import (
	dockerfileGolang "github.com/helmutkemper/iotmaker.docker.builder.golang.dockerfile"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
	"time"
)

// Init
//
// English: Initializes the object and should be called only after all settings have been configured
//
// Português: Inicializa o objeto e deve ser chamado apenas depois de toas as configurações serem definidas
func (e *ContainerBuilder) Init() (err error) {

	e.restartPolicy = iotmakerdocker.KRestartPolicyNo

	if e.autoDockerfile == nil {
		e.autoDockerfile = &dockerfileGolang.DockerfileGolang{}
	}

	if e.environmentVar == nil {
		e.environmentVar = make([]string, 0)
	}

	onStart := make(chan bool, 1)
	e.onContainerReady = &onStart

	onInspect := make(chan bool, 1)
	e.onContainerInspect = &onInspect

	e.changePointer = *iotmakerdocker.NewImagePullStatusChannel()

	e.dockerSys = iotmakerdocker.DockerSystem{}
	err = e.dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	if e.inspectInterval != 0 {
		e.ticker = time.NewTicker(e.inspectInterval)

		go func(e *ContainerBuilder) {
			var err error
			var logs []byte

			for {
				select {
				case <-e.ticker.C:

					if e.containerID == "" {
						e.containerID, err = e.dockerSys.ContainerFindIdByName(e.containerName)
						if err != nil {
							continue
						}
					}

					e.inspect, _ = e.dockerSys.ContainerInspectParsed(e.containerID)
					logs, _ = e.dockerSys.ContainerLogs(e.containerID)
					e.logs = string(logs)
					*e.onContainerInspect <- true
				}
			}

		}(e)
	}
	return
}
