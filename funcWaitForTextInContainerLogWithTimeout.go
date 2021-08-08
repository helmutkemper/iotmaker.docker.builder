package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"log"
	"time"
)

// WaitForTextInContainerLogWithTimeout
//
// English: Wait for the text to appear in the container's default output
//
//   value: searched text
//   timeout: wait timeout
//
// Português: Espera pelo texto aparecer na saída padrão do container
//
//   value: texto procurado
//   timeout: tempo limite de espera
func (e *ContainerBuilder) WaitForTextInContainerLogWithTimeout(value string, timeout time.Duration) (dockerLogs string, err error) {
	var logs []byte
	logs, err = e.dockerSys.ContainerLogsWaitTextWithTimeout(e.containerID, value, timeout, log.Writer())
	if err != nil {
		util.TraceToLog()
	}
	return string(logs), err
}
