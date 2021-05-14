package iotmaker_docker_builder

// RemoveAllByNameContains (english):
//
// RemoveAllByNameContains (português): procuta por redes, volumes, container e imagens que contenham o termo definido
// em "value" no nome e tenta remover os mesmos
func (e *ContainerBuilder) RemoveAllByNameContains(value string) (err error) {
	e.containerID = ""
	err = e.dockerSys.RemoveAllByNameContains(value)
	return
}
