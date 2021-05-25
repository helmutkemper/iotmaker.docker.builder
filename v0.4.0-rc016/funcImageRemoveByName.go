package iotmakerdockerbuilder

// ImageRemoveByName (english):
//
// ImageRemoveByName (português): Remove uma imagem se não houver containers usando a imagem
//   name: nome completo da imagem
//
//     Nota: pare e remova todos os containers usando a imagem antes de usar este comando
//
func (e *ContainerBuilder) ImageRemoveByName(name string) (err error) {
	err = e.dockerSys.ImageRemoveByName(name, false, false)
	if err != nil {
		return
	}

	return
}
