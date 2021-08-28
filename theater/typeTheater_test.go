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

	var containerServer = buildServerContainer(t)

	var server = NewContainer(&containerServer).
		SetLogPath("./test.log.csv").
		AddFilterToLog("contador", "blabla", "^.*?counter: (?P<valueToGet>[\\d\\.]+)", "\\.", ",").
		AddCaosPauseDuration(time.Second, 3*time.Second).
		AddCaosUnpauseDuration(time.Second, 3*time.Second).
		AddCaosStartDuration(time.Second, 3*time.Second)

	var theater = Theater{}
	err = theater.AddContainers(server)
	if err != nil {
		util.TraceToLog()
		log.Printf("err: %v", err)
		t.Fail()
	}

	err = theater.AddContainerToCaosScene(server)
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

func buildServerContainer(t *testing.T) (container builder.ContainerBuilder) {
	var err error

	container = builder.ContainerBuilder{}

	container.SetPrintBuildOnStrOut()
	container.SetCacheEnable(true)
	container.MakeDefaultDockerfileForMe()
	container.SetImageCacheName("cache:latest")
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
	// replace container folder /static to host folder ./test/static
	err = container.AddFileOrFolderToLinkBetweenConputerHostAndContainer("../test/static", "/static")
	if err != nil {
		util.TraceToLog()
		log.Printf("err: %v", err)
		t.Fail()
	}

	return
}
