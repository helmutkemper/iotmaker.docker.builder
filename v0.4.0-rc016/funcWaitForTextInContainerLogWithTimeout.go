package iotmakerdockerbuilder

import (
	"log"
	"time"
)

// WaitForTextInContainerLogWithTimeout (english):
//
// WaitForTextInContainerLogWithTimeout (português):
func (e *ContainerBuilder) WaitForTextInContainerLogWithTimeout(value string, timeout time.Duration) (dockerLogs string, err error) {
	var logs []byte
	logs, err = e.dockerSys.ContainerLogsWaitTextWithTimeout(e.containerID, value, timeout, log.Writer())
	return string(logs), err
}
