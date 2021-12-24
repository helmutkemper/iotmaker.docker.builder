package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
)

// GarbageCollector
//
// English:
//
//  Removes docker elements created during unit tests, such as networks, containers, images and volumes with the term delete in the name.
//
//  Eg: network_to_delete_after_test
//
//   Input:
//     names: Terms contained in the name of docker elements indicated for removal.
//       Eg: nats, removes network, container image, and volume elements that contain the term "nats" in the name. [optional]
//
// Português:
//
//  Remove elementos docker criados dutente os testtes unitários, como por exemplo, redes, contêineres, imagens e volumes com o termo delete no nome.
//
//  ex.: network_to_delete_after_test
//
//   Entrada:
//     names: Termos contidos no nome dos elementos docker indicados para remoção.
//       Ex.: nats, remove os elementos de rede, imagem container e volumes que contenham o termo "nats" no nome. [opcional]
func GarbageCollector(names ...string) {
	var err error

	// garbage collector delete all containers, images, volumes and networks whose name contains the term "delete"
	var garbageCollector = ContainerBuilder{}
	err = garbageCollector.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	// set the term "delete" to garbage collector
	err = garbageCollector.RemoveAllByNameContains("delete")
	if err != nil {
		util.TraceToLog()
		return
	}

	for _, nameContains := range names {
		_ = garbageCollector.RemoveAllByNameContains(nameContains)
	}
}
