package iotmakerdockerbuilder

// SetImageName
//
// English: Defines the name of the image to be used or created
//
//   value: name of the image to be downloaded or created
//
//     Note: the image name must have the version tag
//
// Português: Define o nome da imagem a ser usada ou criada
//
//   value: noma da imagem a ser baixada ou criada
//
//     Nota: o nome da imagem deve ter a tag de versão
func (e *ContainerBuilder) SetImageName(value string) {
	e.imageName = value
}
