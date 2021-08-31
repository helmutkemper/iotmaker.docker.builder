package theater

import (
	"bytes"
	builder "github.com/helmutkemper/iotmaker.docker.builder"
	"github.com/helmutkemper/util"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

func TestTheater_BuildFromImage(t *testing.T) {
	var err error

	builder.GarbageCollector()

	// create a container
	var containerNats = builder.ContainerBuilder{}
	// imprime a saída padrão do container
	containerNats.SetPrintBuildOnStrOut()
	// set image name for docker pull
	containerNats.SetImageName("nats:latest")
	// set a container name
	containerNats.SetContainerName("nats_delete_after_test")
	// set a waits for the text to appear in the standard container output to proceed [optional]
	containerNats.SetWaitStringWithTimeout("Listening for route connections on 0.0.0.0:6222", 10*time.Second)

	var containerNatsConfiguration = NewTestContainerConfiguration(&containerNats).
		SetASceneLinearFlag()

	// Theater
	var theater = Theater{}

	// Add first scene
	err = theater.AddContainerConfiguration(containerNatsConfiguration)
	if err != nil {
		util.TraceToLog()
		log.Printf("err: %v", err)
		t.Fail()
	}

	// Init theater
	err = theater.Init()
	if err != nil {
		util.TraceToLog()
		log.Printf("err: %v", err)
		t.Fail()
	}

	builder.GarbageCollector()
}

func TestTheater_WriteStatsCSV(t *testing.T) {
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
	container.SetBuildFolderPath("../test/counter")
	// container name container_delete_server_after_test
	container.SetContainerName("container_delete_counter_after_test")
	// define o limite de memória
	container.SetImageBuildOptionsMemory(100 * builder.KMegaByte)

	var containerServerConfiguration = NewTestContainerConfiguration(&container).
		SetContainerStatsLogPath("./counter.log.csv").
		AddASceneSettingFilterOnTheContainersStandardOutputToEndEvent(
			"done!",
			"^.*?(?P<valueToGet>\\d+/\\d+/\\d+ \\d+:\\d+:\\d+ done!).*",
			"",
			"",
		).
		AddASceneSettingFilterOnTheContainersStandardOutputToFailEvent(
			"counter: 40",
			"^.*?(?P<valueToGet>\\d+/\\d+/\\d+ \\d+:\\d+:\\d+ counter: [\\d\\.]+).*",
			"(?P<date>\\d+/\\d+/\\d+)\\s+(?P<hour>\\d+:\\d+:\\d+)\\s+counter:\\s+(?P<value>[\\d\\.]+).*",
			"Test Fail! Counter Value: ${value} - Hour: ${hour} - Date: ${date}",
		).
		SetASceneLinearFlag()

	// Theater
	var theater = Theater{}

	// Add second scene
	err = theater.AddContainerConfiguration(containerServerConfiguration)
	if err != nil {
		util.TraceToLog()
		log.Printf("err: %v", err)
		t.Fail()
	}

	// Init theater
	err = theater.Init()
	if err != nil {
		util.TraceToLog()
		log.Printf("err: %v", err)
		t.Fail()
	}

	theater.SetLogFields(KMacOsLog)

	eventCh := theater.GetChannels()

	select {
	case event := <-*eventCh:
		if event.Error == true {
			t.Fail()
			log.Printf("error: %v", event.Message)
			break
		}

		if event.Done == true {
			log.Printf("test pass: %v", event.Message)
			break
		}

		if event.Fail == true {
			t.Fail()
			log.Printf("test fail: %v", event.Message)
			break
		}
	}

	builder.GarbageCollector()

	var file []byte
	file, err = ioutil.ReadFile("./counter.log.csv")
	if err != nil {
		t.Fail()
		log.Printf("error: %v", err.Error())
	}

	var counter = 0
	var line = make([]byte, 0)
	for _, char := range file {
		if char == []byte("\n")[0] {
			counter += 1
			if counter == 1 {
				var lineReader = []byte(
					"Reading time\t" +
						"Linux specific stats, not populated on Windows. Current is the number of pids in the cgroup\t" +
						"Total CPU time consumed. (Units: nanoseconds on Linux, Units: 100's of nanoseconds on Windows)\t" +
						"Time spent by tasks of the cgroup in kernel mode (Units: nanoseconds on Linux). Time spent by all container processes in kernel mode (Units: 100's of nanoseconds on Windows.Not populated for Hyper-V Containers.).\t" +
						"System Usage. Linux only.\t" +
						"Online CPUs. Linux only.\t" +
						"Total CPU time consumed. (Units: nanoseconds on Linux. Units: 100's of nanoseconds on Windows)\t" +
						"Time spent by tasks of the cgroup in kernel mode (Units: nanoseconds on Linux) - Time spent by all container processes in kernel mode (Units: 100's of nanoseconds on Windows - Not populated for Hyper-V Containers.)\t" +
						"System Usage. (Linux only)\t" +
						"Online CPUs. (Linux only)\t" +
						"Current res_counter usage for memory\t" +
						"Maximum usage ever recorded.\t" +
						"Memory limit\t",
				)
				if bytes.Equal(line, lineReader) == false {
					t.Fail()
					log.Printf("CSV reader fail")
				}
			}
		} else {
			line = append(line, char)
		}
	}

	//_ = os.Remove("./counter.log.csv")

}

func TestTheater_AddContainers(t *testing.T) {
	var err error

	builder.GarbageCollector()

	// *******************************************************************************************************************
	// Nats - container
	// *******************************************************************************************************************

	// create a container
	var containerNats = builder.ContainerBuilder{}
	// imprime a saída padrão do container
	containerNats.SetPrintBuildOnStrOut()
	// set image name for docker pull
	containerNats.SetImageName("nats:latest")
	// set a container name
	containerNats.SetContainerName("nats_delete_after_test")
	// set a waits for the text to appear in the standard container output to proceed [optional]
	containerNats.SetWaitStringWithTimeout("Listening for route connections on 0.0.0.0:6222", 10*time.Second)

	// *******************************************************************************************************************
	// Nats - scene configuration
	// *******************************************************************************************************************

	var containerNatsConfiguration = NewTestContainerConfiguration(&containerNats).
		SetASceneLinearFlag()

	// *******************************************************************************************************************
	// Project in folder - container
	// *******************************************************************************************************************

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
	container.SetBuildFolderPath("../test/counter")
	// container name container_delete_server_after_test
	container.SetContainerName("container_delete_server_after_test")
	// define o limite de memória
	container.SetImageBuildOptionsMemory(100 * builder.KMegaByte)

	// replace container folder /static to host folder ./test/static
	err = container.AddFileOrFolderToLinkBetweenConputerHostAndContainer("../test/static", "/static")
	if err != nil {
		util.TraceToLog()
		log.Printf("err: %v", err)
		t.Fail()
	}

	// *******************************************************************************************************************
	// Project in folder - scene configuration
	// *******************************************************************************************************************

	var containerServerConfiguration = NewTestContainerConfiguration(&container).
		SetContainerStatsLogPath("./counter.log.csv").
		//AddFilterToCaptureInformationOnTheContainersStandardOutputForStatsLog("contador", "blabla", "^.*?counter: (?P<valueToGet>[\\d\\.]+)", "\\.", ",").
		//AddASceneOfChaosSettingPauseDuration(5*time.Second, 10*time.Second).
		//AddASceneOfChaosSettingUnpauseDuration(20*time.Second, 30*time.Second).
		//AddASceneOfChaosSettingStartDuration(5*time.Second, 15*time.Second).
		//AddASceneOfChaosSettingRestartInterval(20*time.Second, 30*time.Second).
		//AddASceneOfChaosSettingRestartIntervalRestartController(0.2, -1).
		AddASceneSettingFilterOnTheContainersStandardOutputToEndEvent("done!", "^.*?(?P<valueToGet>\\d+/\\d+/\\d+ \\d+:\\d+:\\d+ done!).*", "", "").
		//AddASceneSettingFilterOnTheContainersStandardOutputToFailEvent(
		//	"counter: 20",
		//	"^.*?(?P<valueToGet>\\d+/\\d+/\\d+ \\d+:\\d+:\\d+ counter: [\\d\\.]+).*",
		//	"",
		//	"",
		//).
		//AddASceneSettingFilterOnTheContainersStandardOutputToFailEvent(
		//	"counter: 20",
		//	"^.*?(?P<valueToGet>\\d+/\\d+/\\d+ \\d+:\\d+:\\d+ counter: [\\d\\.]+).*",
		//	"(\\d+/\\d+/\\d+)\\s+(\\d+:\\d+:\\d+)\\s+counter:\\s+([\\d\\.]+).*",
		//	"Value: $3 - $2 - $1",
		//).
		AddASceneSettingFilterOnTheContainersStandardOutputToFailEvent(
			"counter: 20",
			"^.*?(?P<valueToGet>\\d+/\\d+/\\d+ \\d+:\\d+:\\d+ counter: [\\d\\.]+).*",
			"(?P<date>\\d+/\\d+/\\d+)\\s+(?P<hour>\\d+:\\d+:\\d+)\\s+counter:\\s+(?P<value>[\\d\\.]+).*",
			"Value: ${value} - Hour: ${hour} - Date: ${date}",
		).
		SetASceneLinearFlag()

	// Theater
	var theater = Theater{}

	// Add first scene
	err = theater.AddContainerConfiguration(containerNatsConfiguration)
	if err != nil {
		util.TraceToLog()
		log.Printf("err: %v", err)
		t.Fail()
	}

	// Add second scene
	err = theater.AddContainerConfiguration(containerServerConfiguration)
	if err != nil {
		util.TraceToLog()
		log.Printf("err: %v", err)
		t.Fail()
	}

	// Init theater
	err = theater.Init()
	if err != nil {
		util.TraceToLog()
		log.Printf("err: %v", err)
		t.Fail()
	}

	eventCh := theater.GetChannels()

	for {
		select {
		case e := <-*eventCh:
			log.Printf("Event: %+v", e)
		}
	}

}
