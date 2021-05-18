package iotmaker_docker_builder

// GetImageID (english):
//
// GetImageID (português): Retorna o ID da imagem.
func (e *ContainerBuilder) GetImageID() (ID string) {
	return e.imageID
}
