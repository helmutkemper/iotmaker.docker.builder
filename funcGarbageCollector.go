package iotmaker_docker_builder

import (
	"github.com/helmutkemper/util"
)

// GarbageCollector (englis): All tests use networks, containers and images with the term "delete" contained in the
// name.
//
// This function considers that the test is over and that these elements must be removed at the end of each test, and as
// a guarantee, if any test has failed, it is also used before each test.
//
// GarbageCollector (português): Todos os testes usam redes, containers e imagens com o termo "delete" contido no nome.
//
// Esta função considera que o teste acabou e que estes elementos devem ser removidos ao final de cada teste, e por
// garantia, caso algum teste tenha falhado, também é usada antes de cada teste.
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
