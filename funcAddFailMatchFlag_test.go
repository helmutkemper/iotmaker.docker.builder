package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"log"
	"time"
)

func ExampleContainerBuilder_AddFailMatchFlag() {
	var err error
	var imageInspect types.ImageInspect

	GarbageCollector()

	var container = ContainerBuilder{}
	// imprime a saída padrão do container
	container.SetPrintBuildOnStrOut()
	// caso exista uma imagem de nome cache:latest, ela será usada como base para criar o container
	container.SetCacheEnable(true)
	// monta um dockerfile padrão para o golang onde o arquivo main.go e o arquivo go.mod devem está na pasta raiz
	container.MakeDefaultDockerfileForMe()
	// new image name delete:latest
	container.SetImageName("delete:latest")
	// set a folder path to make a new image
	container.SetBuildFolderPath("./test/counter")
	// container name container_delete_server_after_test
	container.SetContainerName("container_counter_delete_after_test")
	// define o limite de memória
	container.SetImageBuildOptionsMemory(100 * KMegaByte)

	container.AddFailMatchFlag(
		"counter: 40",
	)

	err = container.Init()
	if err != nil {
		fmt.Printf("error: %v", err.Error())
		GarbageCollector()
		return
	}

	imageInspect, err = container.ImageBuildFromFolder()
	if err != nil {
		fmt.Printf("error: %v", err.Error())
		GarbageCollector()
		return
	}

	fmt.Printf("image size: %v\n", container.SizeToString(imageInspect.Size))
	fmt.Printf("image os: %v\n", imageInspect.Os)

	err = container.ContainerBuildAndStartFromImage()
	if err != nil {
		log.Printf("error: %v", err.Error())
		GarbageCollector()
		return
	}

	container.StartMonitor(time.NewTicker(2 * time.Second))

	event := container.GetChaosEvent()

	select {
	case e := <-*event:
		fmt.Printf("container name: %v\n", e.ContainerName)
		fmt.Printf("done: %v\n", e.Done)
		fmt.Printf("fail: %v\n", e.Fail)
		fmt.Printf("error: %v\n", e.Error)
	}

	container.StopMonitor()

	GarbageCollector()

	// Output:
	// image size: 1.38 MB
	// image os: linux
	// container name: container_counter_delete_after_test
	// done: false
	// fail: true
	// error: false
}
