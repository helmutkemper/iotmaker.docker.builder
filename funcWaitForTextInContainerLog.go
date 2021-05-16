package iotmaker_docker_builder

import (
	"log"
	"time"
)

// WaitForTextInContainerLog (english):
//
// WaitForTextInContainerLog (português): Para a execução do objeto até o texto ser encontrado na saída padrão do
// container
//   value: texto indicativo de evento apresentado na saída padrão do container
func (e *ContainerBuilder) WaitForTextInContainerLog(value string) (dockerLogs string, err error) {
	var logs []byte
	logs, err = e.dockerSys.ContainerLogsWaitText(e.containerID, value, log.Writer())
	return string(logs), err
}

func (e *ContainerBuilder) WaitForTextInContainerLogWithTimeout(value string, timeout time.Duration) (dockerLogs string, err error) {
	var logs []byte
	logs, err = e.dockerSys.ContainerLogsWaitTextWithTimeout(e.containerID, value, timeout, log.Writer())
	return string(logs), err
}
