package theater

import (
	builder "github.com/helmutkemper/iotmaker.docker.builder"
	"github.com/helmutkemper/util"
	"log"
	"sync"
	"testing"
	"time"
)

func TestTheater_AddContainers(t *testing.T) {
	var err error

	builder.GarbageCollector()

	var container = builder.ContainerBuilder{}

	// imprime a saída padrão do container
	container.SetPrintBuildOnStrOut()

	// caso exista uma imagem de nome cache:latest, ela será usada como base para criar o container
	container.SetCacheEnable(true)

	// monta um dockerfile padrão para o golang onde o arquivo main.go e o arquivo go.mod devem está na pasta raiz
	container.MakeDefaultDockerfileForMe()

	// new image name delete:latest
	container.SetImageName("delete:latest")

	// set a folder path to make a new image
	container.SetBuildFolderPath("../test/server")

	// container name container_delete_server_after_test
	container.SetContainerName("container_delete_server_after_test")

	// set a waits for the text to appear in the standard container output to proceed [optional]
	container.SetWaitStringWithTimeout("starting server at port", 10*time.Second)

	// change and open port 3000 to 3030
	container.AddPortToExpose("3000")

	// define o limite de memória
	container.SetImageBuildOptionsMemory(500 * builder.KMegaByte)

	// replace container folder /static to host folder ./test/static
	err = container.AddFileOrFolderToLinkBetweenConputerHostAndContainer("../test/static", "/static")
	if err != nil {
		util.TraceToLog()
		log.Printf("err: %v", err)
		t.Fail()
	}

	var containerServerConfiguration = NewTestContainerConfiguration(&container).
		SetContainerStatsLogPath("./test.log.csv").
		AddFilterToCaptureInformationOnTheContainersStandardOutputForStatsLog("contador", "blabla", "^.*?counter: (?P<valueToGet>[\\d\\.]+)", "\\.", ",").
		//SetASceneLinearFlag()
		AddASineOfChaosSettingPauseDuration(time.Second, 3*time.Second).
		AddASineOfChaosSettingUnpauseDuration(time.Second, 3*time.Second).
		AddASineOfChaosSettingStartDuration(time.Second, 3*time.Second).
		AddASineOfChaosSettingRestartInterval(1*time.Second, 2*time.Second).
		AddASineOfChaosSettingRestartIntervalRestartController(0.5, 2)

	var theater = Theater{}
	err = theater.AddContainerConfiguration(containerServerConfiguration)
	if err != nil {
		util.TraceToLog()
		log.Printf("err: %v", err)
		t.Fail()
	}

	err = theater.Init()
	if err != nil {
		util.TraceToLog()
		log.Printf("err: %v", err)
		t.Fail()
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
