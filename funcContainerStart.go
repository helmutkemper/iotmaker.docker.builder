package iotmaker_docker_builder

// ContainerStart (english):
//
// ContainerStart (português): inicializar um container recem criado ou pausado
func (e *ContainerBuilder) ContainerStart() (err error) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			return
		}
	}

	err = e.dockerSys.ContainerStart(e.containerID)
	return
}
