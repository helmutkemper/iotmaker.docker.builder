package iotmaker_docker_builder

// ImagePull (english):
//
// ImagePull (português): baixa a imagem a ser montada. (equivale ao comando docker pull)
func (e *ContainerBuilder) ImagePull() (err error) {
	e.imageID, e.imageName, err = e.dockerSys.ImagePull(e.imageName, e.changePointer)
	if err != nil {
		return
	}

	return
}
