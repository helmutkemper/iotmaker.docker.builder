package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
)

// GarbageCollector
//
// English: A great use of this code is to build container during unit testing, and in this case, you can add the
// term delete to the name of all docker elements created during the test, so that they are deleted in a simple way.
// e.g..: network_to_delete_after_test
//
// Português: Uma grande utilidade desse código é levantar container durante testes unitários, e nesse caso, você
// pode adicionar o termo delete ao nome de todos os elementos docker criado durante o teste, para que os mesmos
// sejam apagados de forma simples.
// ex.: network_to_delete_after_test
func GarbageCollector() {
	var err error

	// garbage collector delete all containers, images, volumes and networks whose name contains the term "delete"
	var garbageCollector = ContainerBuilder{}
	err = garbageCollector.Init()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	// set the term "delete" to garbage collector
	err = garbageCollector.RemoveAllByNameContains("delete")
	if err != nil {
		util.TraceToLog()
		panic(err)
	}
}
