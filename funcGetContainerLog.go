package iotmaker_docker_builder

// GetContainerLog (english):
//
// GetContainerLog (português): baixa a saída padrão do container
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
