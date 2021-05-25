package iotmakerdockerbuilder

// GetContainerLog
//
// English: Returns the current standard output of the container.
//
// Português: Retorna a saída padrão atual do container.
func (e *ContainerBuilder) GetContainerLog() (log []byte, err error) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			return
		}
	}

	log, err = e.dockerSys.ContainerLogs(e.containerID)
	return
}
