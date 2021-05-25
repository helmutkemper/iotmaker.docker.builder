package iotmakerdockerbuilder

// GetImageName
//
// English: Returns the name of the image.
//
// Português: Retorna o nome da imagem.
func (e *ContainerBuilder) GetImageName() (name string) {
	return e.imageName
}
