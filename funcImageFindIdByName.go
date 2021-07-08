package iotmakerdockerbuilder

func (e *ContainerBuilder) ImageFindIdByName(name string) (id string, err error) {
  return e.dockerSys.ImageFindIdByName(name)
}
