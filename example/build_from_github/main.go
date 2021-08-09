package main

import (
	builder "github.com/helmutkemper/iotmaker.docker.builder"
	"time"
)

func main() {
	var err error
	var container = builder.ContainerBuilder{}
	// new image name delete:latest
	container.SetImageName("delete:latest")
	// container name container_delete_server_after_test
	container.SetContainerName("container_delete_server_after_test")
	// git project to clone https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git
	container.SetGitCloneToBuild("https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git")

	// see SetGitCloneToBuildWithUserPassworh(), SetGitCloneToBuildWithPrivateSshKey() and
	// SetGitCloneToBuildWithPrivateToken()

	// set a waits for the text to appear in the standard container output to proceed [optional]
	container.SetWaitStringWithTimeout("Stating server on port 3000", 10*time.Second)
	// change and open port 3000 to 3030
	container.AddPortToChange("3000", "3030")
	// replace container folder /static to host folder ./test/static
	err = container.AddFileOrFolderToLinkBetweenConputerHostAndContainer("../../test/static", "/static")
	if err != nil {
		panic(err)
	}

	// inicialize container object
	err = container.Init()
	if err != nil {
		panic(err)
	}

	// builder new image from git project
	err = container.ImageBuildFromServer()
	if err != nil {
		panic(err)
	}

	err = container.ContainerBuildAndStartFromImage()
	if err != nil {
		panic(err)
	}
}
