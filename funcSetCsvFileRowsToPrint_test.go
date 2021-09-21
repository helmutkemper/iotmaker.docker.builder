package iotmakerdockerbuilder

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestContainerBuilder_SetCsvFileRowsToPrint_2(t *testing.T) {
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
		"",
		"",
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

	container.SetCsvFileRowsToPrint(KReadingTime | KCurrentNumberOfOidsInTheCGroup)

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

func TestContainerBuilder_SetCsvFileRowsToPrint_3(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.3.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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

	container.SetCsvFileRowsToPrint(KReadingTime | KCurrentNumberOfOidsInTheCGroup | KLimitOnTheNumberOfPidsInTheCGroup)

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

func TestContainerBuilder_SetCsvFileRowsToPrint_4(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.4.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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

	container.SetCsvFileRowsToPrint(KReadingTime | KCurrentNumberOfOidsInTheCGroup | KLimitOnTheNumberOfPidsInTheCGroup | KTotalCPUTimeConsumed)

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

func TestContainerBuilder_SetCsvFileRowsToPrint_5(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.5.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_6(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.6.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_7(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.7.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_8(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.8.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_9(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.9.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_10(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.10.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_11(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.11.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_12(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.12.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_13(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.13.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_14(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.14.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_15(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.15.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_16(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.16.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_17(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.17.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_18(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.18.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_19(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.19.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_20(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.20.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_21(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.21.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_22(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.22.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_23(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.23.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_24(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.24.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KNumberOfTimesMemoryUsageHitsLimits,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_25(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.25.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KNumberOfTimesMemoryUsageHitsLimits |
		KMemoryLimit,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_26(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.26.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KNumberOfTimesMemoryUsageHitsLimits |
		KMemoryLimit |
		KCommittedBytes,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_27(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.27.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KNumberOfTimesMemoryUsageHitsLimits |
		KMemoryLimit |
		KCommittedBytes |
		KPeakCommittedBytes,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_28(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.28.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KNumberOfTimesMemoryUsageHitsLimits |
		KMemoryLimit |
		KCommittedBytes |
		KPeakCommittedBytes |
		KPrivateWorkingSet,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_29(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.29.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KNumberOfTimesMemoryUsageHitsLimits |
		KMemoryLimit |
		KCommittedBytes |
		KPeakCommittedBytes |
		KPrivateWorkingSet |
		KBlkioIoServiceBytesRecursive,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_30(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.30.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KNumberOfTimesMemoryUsageHitsLimits |
		KMemoryLimit |
		KCommittedBytes |
		KPeakCommittedBytes |
		KPrivateWorkingSet |
		KBlkioIoServiceBytesRecursive |
		KBlkioIoServicedRecursive,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_31(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.31.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KNumberOfTimesMemoryUsageHitsLimits |
		KMemoryLimit |
		KCommittedBytes |
		KPeakCommittedBytes |
		KPrivateWorkingSet |
		KBlkioIoServiceBytesRecursive |
		KBlkioIoServicedRecursive |
		KBlkioIoQueuedRecursive,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_32(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.32.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KNumberOfTimesMemoryUsageHitsLimits |
		KMemoryLimit |
		KCommittedBytes |
		KPeakCommittedBytes |
		KPrivateWorkingSet |
		KBlkioIoServiceBytesRecursive |
		KBlkioIoServicedRecursive |
		KBlkioIoQueuedRecursive |
		KBlkioIoServiceTimeRecursive,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_33(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.33.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KNumberOfTimesMemoryUsageHitsLimits |
		KMemoryLimit |
		KCommittedBytes |
		KPeakCommittedBytes |
		KPrivateWorkingSet |
		KBlkioIoServiceBytesRecursive |
		KBlkioIoServicedRecursive |
		KBlkioIoQueuedRecursive |
		KBlkioIoServiceTimeRecursive |
		KBlkioIoWaitTimeRecursive,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_34(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.34.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KNumberOfTimesMemoryUsageHitsLimits |
		KMemoryLimit |
		KCommittedBytes |
		KPeakCommittedBytes |
		KPrivateWorkingSet |
		KBlkioIoServiceBytesRecursive |
		KBlkioIoServicedRecursive |
		KBlkioIoQueuedRecursive |
		KBlkioIoServiceTimeRecursive |
		KBlkioIoWaitTimeRecursive |
		KBlkioIoMergedRecursive,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_35(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.35.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KNumberOfTimesMemoryUsageHitsLimits |
		KMemoryLimit |
		KCommittedBytes |
		KPeakCommittedBytes |
		KPrivateWorkingSet |
		KBlkioIoServiceBytesRecursive |
		KBlkioIoServicedRecursive |
		KBlkioIoQueuedRecursive |
		KBlkioIoServiceTimeRecursive |
		KBlkioIoWaitTimeRecursive |
		KBlkioIoMergedRecursive |
		KBlkioIoTimeRecursive,
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

func TestContainerBuilder_SetCsvFileRowsToPrint_36(t *testing.T) {
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

	container.SetLogPath("./test.counter.log.36.csv")
	container.AddFilterToLog(
		"contador",
		"counter",
		"^.*?counter: (?P<valueToGet>[\\d\\.]+)",
		"",
		"",
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
		KCurrentNumberOfOidsInTheCGroup |
		KLimitOnTheNumberOfPidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KNumberOfPeriodsWithThrottlingActive |
		KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit |
		KAggregateTimeTheContainerWasThrottledForInNanoseconds |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KAggregatePreCPUTimeTheContainerWasThrottled |
		KNumberOfPeriodsWithPreCPUThrottlingActive |
		KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KNumberOfTimesMemoryUsageHitsLimits |
		KMemoryLimit |
		KCommittedBytes |
		KPeakCommittedBytes |
		KPrivateWorkingSet |
		KBlkioIoServiceBytesRecursive |
		KBlkioIoServicedRecursive |
		KBlkioIoQueuedRecursive |
		KBlkioIoServiceTimeRecursive |
		KBlkioIoWaitTimeRecursive |
		KBlkioIoMergedRecursive |
		KBlkioIoTimeRecursive |
		KBlkioSectorsRecursive,
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
