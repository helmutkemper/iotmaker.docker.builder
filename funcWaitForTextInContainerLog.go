package iotmakerdockerbuilder

import (
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
	return string(logs), err
}
