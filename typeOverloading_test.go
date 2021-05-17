package iotmaker_docker_builder

import (
	"github.com/helmutkemper/util"
	"log"
	"time"
)

func ExampleOverloading_Init() {
	var err error

	GarbageCollector()

	var serverContainer = &ContainerBuilder{}
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
	serverContainer.AddPortToChange("3000", "3030")
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

	var overload = Overloading{}
	overload.SetBuilderToOverload(serverContainer)
	err = overload.Init()
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		panic(err)
	}

	GarbageCollector()

	// Output:
	//
}
