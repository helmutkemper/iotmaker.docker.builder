package iotmakerdockerbuilder

import (
	dockerNetwork "github.com/helmutkemper/iotmaker.docker.builder.network"
	"github.com/helmutkemper/util"
	"log"
	"time"
)

func ExampleOverloading_Init() {
	var err error

	GarbageCollector()

	var mongoDocker = &ContainerBuilder{}
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
	return
	var netDocker = dockerNetwork.ContainerBuilderNetwork{}
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

	var serverContainer = &ContainerBuilder{}
	// add container to network
	serverContainer.SetNetworkDocker(&netDocker)
	// new image name delete:latest
	serverContainer.SetImageName("delete:latest")
	// container name container_delete_server_after_test
	serverContainer.SetContainerName("container_delete_server_after_test")
	// git project to clone https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git
	serverContainer.SetGitCloneToBuild("https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git")

	// see SetGitCloneToBuildWithUserPassworh(), SetGitCloneToBuildWithPrivateSshKey() and
	// SetGitCloneToBuildWithPrivateToken()

	// set a waits for the text to appear in the standard container output to proceed [optional]
	serverContainer.SetWaitStringWithTimeout("Stating server on port 3000", 10*time.Second)
	// change and open port 3000 to 3030
	//serverContainer.SetDoNotOpenContainersPorts()
	//serverContainer.AddPortToChange("3000", "3030")
	serverContainer.AddPortToExpose("3000")
	// replace container folder /static to host folder ./test/static
	err = serverContainer.AddFiileOrFolderToLinkBetweenConputerHostAndContainer("./test/static", "/static")
	if err != nil {
		log.Printf("container.AddFiileOrFolderToLinkBetweenConputerHostAndContainer().error: %v", err.Error())
		util.TraceToLog()
		panic(err)
	}

	// inicialize container object
	err = serverContainer.Init()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	// builder new image from git project
	err = serverContainer.ImageBuildFromServer()
	if err != nil {
		util.TraceToLog()
		log.Printf("container.ImageBuildFromServer().error: %v", err.Error())
		panic(err)
	}

	// container build from image delete:latest
	err = serverContainer.ContainerBuildFromImage()
	if err != nil {
		util.TraceToLog()
		log.Printf("container.ContainerBuildFromImage().error: %v", err.Error())
		panic(err)
	}

	log.Printf("ip address: %v", serverContainer.GetIPV4Address())
	return
	var overload = Overloading{}
	overload.SetNetworkDocker(&netDocker)
	overload.SetBuilderToOverload(serverContainer)
	err = overload.Init("3000", false)
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		panic(err)
	}

	//Output:
	//
}
