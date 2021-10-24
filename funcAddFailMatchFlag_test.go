package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"io/ioutil"
	"log"
	"os"
	"time"
)

//kemper aqui

func ExampleContainerBuilder_AddFailMatchFlag() {
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

	// English: Golang project path to be turned into docker image
	// Português: Caminho do projeto em Golang a ser transformado em imagem docker
	container.SetBuildFolderPath("./test/bug")

	// English: Defines the name of the docker container to be created.
	// Português: Define o nome do container docker a ser criado.
	container.SetContainerName("container_counter_delete_after_test")

	// English: Defines the maximum amount of memory to be used by the docker container.
	// Português: Define a quantidade máxima de memória a ser usada pelo container docker.
	container.SetImageBuildOptionsMemory(100 * KMegaByte)

	// Português: Define um log, na forma de arquivo CSV, de desempenho do container, com indicadores de consumo de memória e tempos de acesso. Nota: O formato do log varia de acordo com a plataforma, macos, windows, linux.
	// English: Defines a log, in the form of a CSV file, of the container's performance, with indicators of memory consumption and access times. Note: The log format varies by platform, macos, windows, linux.
	container.SetLogPath("./test.counter.log.csv")

	// Português: Define um filtro de busca por texto na saída padrão do container e escreve o texto no log definido por SetLogPath()
	// English: Sets a text search filter on the container's standard output and writes the text to the log defined by SetLogPath()
	container.AddFilterAndReplaceToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"\\.",
		",",
	)

	// English: Adds a failure indicator to the project. Failure indicator is a text searched for in the container's standard output and indicates something that should not have happened during the test.
	// Português: Adiciona um indicador de falha ao projeto. Indicador de falha é um texto procurado na saída padrão do container e indica algo que não deveria ter acontecido durante o teste.
	container.AddFailMatchFlag(
		"counter: 40",
	)

	// English: Adds a log file write failure indicator to the project. Failure indicator is a text searched for in the container's standard output and indicates something that should not have happened during the test.
	// Português: Adiciona um indicador de falha com gravação de arquivo em log ao projeto. Indicador de falha é um texto procurado na saída padrão do container e indica algo que não deveria ter acontecido durante o teste.
	err = container.AddFailMatchFlagToFileLog(
		"bug:",
		"./log1/log2/log3",
	)
	if err != nil {
		fmt.Printf("error: %v", err.Error())
		GarbageCollector()
		return
	}

	// English: Initializes the container manager object.
	// Português: Inicializa o objeto gerenciador de container.
	err = container.Init()
	if err != nil {
		fmt.Printf("error: %v", err.Error())
		GarbageCollector()
		return
	}

	// English: Creates an image from a project folder.
	// Português: Cria uma imagem a partir de uma pasta de projeto.
	imageInspect, err = container.ImageBuildFromFolder()
	if err != nil {
		fmt.Printf("error: %v", err.Error())
		GarbageCollector()
		return
	}

	fmt.Printf("image size: %v\n", container.SizeToString(imageInspect.Size))
	fmt.Printf("image os: %v\n", imageInspect.Os)

	// English: Creates and initializes the container based on the created image.
	// Português: Cria e inicializa o container baseado na imagem criada.
	err = container.ContainerBuildAndStartFromImage()
	if err != nil {
		log.Printf("error: %v", err.Error())
		GarbageCollector()
		return
	}

	// English: Starts container monitoring at two second intervals. This functionality monitors the container's standard output and generates the log defined by the SetLogPath() function.
	// Português: Inicializa o monitoramento do container com intervalos de dois segundos. Esta funcionalidade monitora a saída padrão do container e gera o log definido pela função SetLogPath().
	container.StartMonitor(time.NewTicker(2 * time.Second))

	// English: Gets the event channel inside the container.
	// Português: Pega o canal de eventos dentro do container.
	event := container.GetChaosEvent()

	var fail bool
	for {
		select {
		case e := <-event:
			if e.Fail == true {
				fmt.Printf("container name: %v\n", e.ContainerName)
				fmt.Printf("done: %v\n", e.Done)
				fmt.Printf("fail: %v\n", e.Fail)
				fmt.Printf("error: %v\n", e.Error)

				fail = e.Fail
			}
		}

		if fail == true {
			break
		}
	}

	// English: For container monitoring. Note: This function should be used to avoid trying to read a container that no longer exists, erased by the GarbageCollector() function.
	// Português: Para o monitoramento do container. Nota: Esta função deve ser usada para evitar tentativa de leitura em um container que não existe mais, apagado pela função GarbageCollector().
	_ = container.StopMonitor()

	// English: Deletes all docker elements with the term `delete` in the name.
	// Português: Apaga todos os elementos docker com o termo `delete` no nome.
	GarbageCollector()

	var data []byte
	data, err = ioutil.ReadFile("./log1/log2/log3/log.0.log")
	if err != nil {
		log.Printf("error: %v", err.Error())
		GarbageCollector()
		return
	}

	if len(data) == 0 {
		fmt.Println("log file error")
	}

	_ = os.Remove("./log1/log2/log3/log.0.log")
	_ = os.Remove("./log1/log2/log3/")
	_ = os.Remove("./log1/log2/")
	_ = os.Remove("./log1/")

	// Output:
	// image size: 1.4 MB
	// image os: linux
	// container name: container_counter_delete_after_test
	// done: false
	// fail: true
	// error: false
}
