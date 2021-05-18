package iotmakerdockerbuilder

// ImageRemove (english):
//
// ImageRemove (português): remove a imagem se não houver containers usando a imagem (remova todos os containers antes
// do uso, mesmo os parados)
func (e *ContainerBuilder) ImageRemove() (err error) {
	err = e.dockerSys.ImageRemoveByName(e.imageName, false, false)
	return
}
