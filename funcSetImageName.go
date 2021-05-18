package iotmakerdockerbuilder

// SetImageName (english):
//
// SetImageName (português): Define o nome da imagem a ser usada ou criada
//   value: noma da imagem a ser baixada ou criada
//
//     Nota: o nome da imagem deve ter a tag de versão
func (e *ContainerBuilder) SetImageName(value string) {
	e.imageName = value
}
