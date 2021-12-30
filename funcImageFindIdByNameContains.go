package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

// ImageFindIdByNameContains
//
// English:
//
//  Find an image by part of the name
//
//   Input:
//     containerName: Part of the name of the image
//
//   Output:
//     list: List of images found
//     err: Default error object
//
// Português:
//
//  Encontra uma imagem por parte do nome
//
//   Entrada:
//     containerName: Parte do nome da imagem
//
//   Saída:
//     list: Lista de imagens encontradas
//     err: Objeto de erro padrão
func (e *ContainerBuilder) ImageFindIdByNameContains(containsName string) (list []NameAndId, err error) {
	list = make([]NameAndId, 0)

	e.dockerSys = iotmakerdocker.DockerSystem{}
	err = e.dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	var recevedLis []iotmakerdocker.NameAndId
	recevedLis, err = e.dockerSys.ImageFindIdByNameContains(containsName)
	if err != nil {
		util.TraceToLog()
		return
	}

	for _, elementInList := range recevedLis {
		list = append(list, NameAndId(elementInList))
	}

	return
}
