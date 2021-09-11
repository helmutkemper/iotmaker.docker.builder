package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"log"
	"time"
)

func ExampleContainerBuilder_AddFilterToFail() {
	var err error
	var imageInspect types.ImageInspect

	// English: Deletes all docker elements with the term `delete` in the name.
	// Português: Apaga todos os elementos docker com o termo `delete` no nome.
	GarbageCollector()

	var container = ContainerBuilder{}

	// English: print the standard output of the container
	// Português: imprime a saída padrão do container
	container.SetPrintBuildOnStrOut()

	// English: If there is an image named `cache:latest`, it will be used as a base to create the container.
	// Português: Caso exista uma imagem de nome `cache:latest`, ela será usada como base para criar o container.
	container.SetCacheEnable(true)

	// English: Mount a default dockerfile for golang where the `main.go` file and the `go.mod` file should be in the root folder
	// Português: Monta um dockerfile padrão para o golang onde o arquivo `main.go` e o arquivo `go.mod` devem está na pasta raiz
	container.MakeDefaultDockerfileForMe()

	// English: Name of the new image to be created.
	// Português: Nome da nova imagem a ser criada.
	container.SetImageName("delete:latest")

	// English: Defines the path where the golang code to be transformed into a docker image is located.
	// Português: Define o caminho onde está o código golang a ser transformado em imagem docker.
	container.SetBuildFolderPath("./test/counter")

	// English: Defines the name of the docker container to be created.
	// Português: Define o nome do container docker a ser criado.
	container.SetContainerName("container_counter_delete_after_test")

	// English: Defines the maximum amount of memory to be used by the docker container.
	// Português: Define a quantidade máxima de memória a ser usada pelo container docker.
	container.SetImageBuildOptionsMemory(100 * KMegaByte)

	// English: Defines the log file path with container statistical data
	// Português: Define o caminho do arquivo de log com dados estatísticos do container
	container.SetLogPath("./test.counter.log.csv")

	// English: Adds a search filter to the standard output of the container, to save the information in the log file
	// Português: Adiciona um filtro de busca na saída padrão do container, para salvar a informação no arquivo de log
	container.AddFilterToLog(
		// English: Label to be written to log file
		// Português: Rótulo a ser escrito no arquivo de log
		"contador",

		// English: Simple text searched in the container's standard output to activate the filter
		// Português: Texto simples procurado na saída padrão do container para ativar o filtro
		"counter",

		// Regular expression used to filter what goes into the log using the `valueToGet` parameter.
		// Expressão regular usada para filtrar o que vai para o log usando o parâmetro `valueToGet`.
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",

		// Regular expression used for search and replacement in the text found in the previous step [optional].
		// Expressão regular usada para busca e substituição no texto encontrado na etapa anterior [opcional].
		"\\.",
		",",
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
		fmt.Printf("message: %v\n", e.Message)
	}

	container.StopMonitor()

	GarbageCollector()

	// Output:
	// image size: 1.38 MB
	// image os: linux
	// container name: container_counter_delete_after_test
	// done: true
	// fail: false
	// error: false
	// message: done!
}
