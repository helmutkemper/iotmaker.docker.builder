package iotmakerdockerbuilder

// SetImageBuildOptionsSquash
//
// English: squash the resulting image's layers to the parent preserves the original image and creates a new one
// from the parent with all the changes applied to a single layer
//
// Português: Usa o conteúdo dos layers da imagem pai para criar uma imagem nova, preservando a imagem pai, e aplica
// todas as mudanças a um novo layer
func (e *ContainerBuilder) SetImageBuildOptionsSquash(value bool) {
	e.buildOptions.Squash = value
}
