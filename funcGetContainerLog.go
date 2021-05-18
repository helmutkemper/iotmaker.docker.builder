package iotmakerdockerbuilder

// GetContainerLog (english):
//
// GetContainerLog (português): Retorna a saída padrão atual do container.
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
