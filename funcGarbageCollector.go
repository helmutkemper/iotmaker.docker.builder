package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
)

// GarbageCollector
//
// English: A great use of this code is to build container during unit testing, and in this case, you can add the
// term delete to the name of all docker elements created during the test, so that they are deleted in a simple way.
// e.g..: network_to_delete_after_test
//   Input:
//     names: text contained in docker element name indicated for removal. Ex.: nats, removes network elements,
//            container, image and volumes that contain the term "nats" in the name. [optional]
//
// Português: Uma grande utilidade desse código é levantar container durante testes unitários, e nesse caso, você
// pode adicionar o termo delete ao nome de todos os elementos docker criado durante o teste, para que os mesmos
// sejam apagados de forma simples.
// ex.: network_to_delete_after_test
//   Entrada:
//     names: Nomes contidos nos elementos docker indicados para remoção. Ex.: nats, remove os elementos de rede, imagem
//            container e volumes que contenham o termo "nats" no nome. [opcional]
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
