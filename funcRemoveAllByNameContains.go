package iotmakerdockerbuilder

// RemoveAllByNameContains
//
// English: searches for networks, volumes, containers and images that contain the term defined in "value" in the
// name, and tries to remove them from docker
//
// Português: procura por redes, volumes, container e imagens que contenham o termo definido em "value" no nome, e
// tenta remover os mesmos do docker
func (e *ContainerBuilder) RemoveAllByNameContains(value string) (err error) {
	e.containerID = ""
	err = e.dockerSys.RemoveAllByNameContains(value)
	return
}
