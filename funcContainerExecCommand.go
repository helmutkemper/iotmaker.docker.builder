package iotmakerdockerbuilder

import "github.com/helmutkemper/util"

// ContainerExecCommand
//
// Português: Executa comandos dentro do container.
//
//   Entrada:
//     commands: lista de comandos. Ex.: ["ls", "-l"]
//
func (e *ContainerBuilder) ContainerExecCommand(
	commands []string,
) (
	exitCode int,
	runing bool,
	stdOutput []byte,
	stdError []byte,
	err error,
) {

	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	exitCode, runing, stdOutput, stdError, err = e.dockerSys.ContainerExecCommand(e.containerID, commands)
	return
}
