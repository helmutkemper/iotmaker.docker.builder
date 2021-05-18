package iotmaker_docker_builder

// GetImageName (english):
//
// GetImageName (português): Retorna o nome da imagem.
func (e *ContainerBuilder) GetImageName() (name string) {
	return e.imageName
}
