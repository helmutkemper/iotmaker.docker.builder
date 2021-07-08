package iotmakerdockerbuilder

func (e *ContainerBuilder) ContainerFindIdByName(name string) (id string, err error) {
  return e.dockerSys.ContainerFindIdByName(name)
}
