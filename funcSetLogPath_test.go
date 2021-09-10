package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"log"
)

func ExampleContainerBuilder_SetLogPath() {
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
	container.SetBuildFolderPath("../test/counter")
	// container name container_delete_server_after_test
	container.SetContainerName("container_delete_counter_after_test")
	// define o limite de memória
	container.SetImageBuildOptionsMemory(100 * KMegaByte)

	container.SetLogPath("./counter.log.csv")
	container.AddFilterToSuccess(
		"done!",
		"^.*?(?P<valueToGet>\\d+/\\d+/\\d+ \\d+:\\d+:\\d+ done!).*",
		"",
		"",
	)
	container.AddFilterToFail(
		"counter: 40",
		"^.*?(?P<valueToGet>\\d+/\\d+/\\d+ \\d+:\\d+:\\d+ counter: [\\d\\.]+).*",
		"(?P<date>\\d+/\\d+/\\d+)\\s+(?P<hour>\\d+:\\d+:\\d+)\\s+counter:\\s+(?P<value>[\\d\\.]+).*",
		"Test Fail! Counter Value: ${value} - Hour: ${hour} - Date: ${date}",
	)

	err = container.Init()
	if err != nil {
		log.Printf("error: %v", err.Error())
		return
	}

	imageInspect, err = container.ImageBuildFromFolder()
	if err != nil {
		log.Printf("error: %v", err.Error())
		return
	}

	fmt.Printf("size: %v", imageInspect.Size)
	fmt.Printf("os: %v", imageInspect.Os)

	err = container.ContainerBuildAndStartFromImage()
	if err != nil {
		log.Printf("error: %v", err.Error())
		return
	}

	event := container.GetChaosEvent()

	select {
	case e := <-*event:
		fmt.Printf("container name: %v", e.ContainerName)
		fmt.Printf("done: %v", e.Done)
		fmt.Printf("fail: %v", e.Fail)
		fmt.Printf("error: %v", e.Error)
		fmt.Printf("message: %v", e.Message)
	}

	// Output:

}
