package iotmakerdockerbuilder

import (
	dockerNetwork "github.com/helmutkemper/iotmaker.docker.builder.network"
	"github.com/helmutkemper/util"
	"time"
)

func ExampleNetworkChaos_Init() {
	var err error

	GarbageCollector()

	var netDocker = &dockerNetwork.ContainerBuilderNetwork{}
	err = netDocker.Init()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	err = netDocker.NetworkCreate("cache_delete_after_test", "10.0.0.0/16", "10.0.0.1")
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	var mongoDocker = &ContainerBuilder{}
	mongoDocker.SetNetworkDocker(netDocker)
	mongoDocker.SetImageName("mongo:latest")
	mongoDocker.SetContainerName("container_delete_mongo_after_test")
	//mongoDocker.AddPortToChange("27017", "27016")
	mongoDocker.AddPortToExpose("27017")
	mongoDocker.SetEnvironmentVar(
		[]string{
			"--host 0.0.0.0",
		},
	)
	mongoDocker.SetWaitStringWithTimeout(`"msg":"Waiting for connections","attr":{"port":27017`, 20*time.Second)
	err = mongoDocker.Init()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	err = mongoDocker.ContainerBuildFromImage()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	var chaos = NetworkChaos{}
	chaos.SetNetworkDocker(netDocker)
	chaos.SetFatherContainer(mongoDocker)
	chaos.SetPorts(20017, 27016, false)
	err = chaos.Init()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	//Output:
	//
}
