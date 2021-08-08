package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"log"
)

// WaitForTextInContainerLog
//
// English: Wait for the text to appear in the container's default output
//
//   value: searched text
//
// Português: Espera pelo texto aparecer na saída padrão do container
//
//   value: texto procurado
func (e *ContainerBuilder) WaitForTextInContainerLog(value string) (dockerLogs string, err error) {
	var logs []byte
	logs, err = e.dockerSys.ContainerLogsWaitText(e.containerID, value, log.Writer())
	if err != nil {
		util.TraceToLog()
	}
	return string(logs), err
}
