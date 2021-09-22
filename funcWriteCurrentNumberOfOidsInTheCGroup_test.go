package iotmakerdockerbuilder

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestContainerBuilder_writeCurrentNumberOfOidsInTheCGroup(t *testing.T) {
	var err error

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

	container.SetLogPath("./test.counter.log.2.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
	)
	container.AddFilterToSuccess(
		"done!",
		"^.*?(?P<valueToGet>\\d+/\\d+/\\d+ \\d+:\\d+:\\d+ done!).*",
		"(?P<date>\\d+/\\d+/\\d+)\\s+(?P<hour>\\d+:\\d+:\\d+)\\s+(?P<value>done!).*",
		"${value}",
	)
	container.AddFilterToFail(
		"counter: 40",
		"^.*?(?P<valueToGet>\\d+/\\d+/\\d+ \\d+:\\d+:\\d+ counter: [\\d\\.]+).*",
		"(?P<date>\\d+/\\d+/\\d+)\\s+(?P<hour>\\d+:\\d+:\\d+)\\s+counter:\\s+(?P<value>[\\d\\.]+).*",
		"Test Fail! Counter Value: ${value} - Hour: ${hour} - Date: ${date}",
	)

	container.SetCsvFileRowsToPrint(KReadingTime |
		KCurrentNumberOfOidsInTheCGroup,
	)

	err = container.Init()
	if err != nil {
		fmt.Printf("error: %v", err.Error())
		GarbageCollector()
		t.Fail()
		return
	}

	_, err = container.ImageBuildFromFolder()
	if err != nil {
		fmt.Printf("error: %v", err.Error())
		GarbageCollector()
		t.Fail()
		return
	}

	err = container.ContainerBuildAndStartFromImage()
	if err != nil {
		log.Printf("error: %v", err.Error())
		GarbageCollector()
		t.Fail()
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
		fmt.Printf("message: %v\n", e.Message)
	}

	err = container.StopMonitor()
	if err != nil {
		log.Printf("error: %v", err.Error())
		GarbageCollector()
		t.Fail()
		return
	}

	GarbageCollector()
}
