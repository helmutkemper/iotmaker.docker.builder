package iotmakerdockerbuilder

import (
  iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

func (e *ContainerBuilder) ImageFindIdByNameContains(containsName string) (list []NameAndId, err error) {
  list = make([]NameAndId, 0)

  var recevedLis []iotmakerdocker.NameAndId
  recevedLis, err = e.dockerSys.ImageFindIdByNameContains(containsName)
  if err != nil {
    return
  }

  for _, elementInList := range recevedLis {
    list = append(list, NameAndId(elementInList))
  }

  return
}
