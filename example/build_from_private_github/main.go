package main

import (
	builder "github.com/helmutkemper/iotmaker.docker.builder"
	"time"
)

// for this project, please, copy the folder server as private project inside your github
func main() {
	var err error

	var container = builder.ContainerBuilder{}
	container.SetPrintBuildOnStrOut()

	// go env -w GOPRIVATE=$GIT_PRIVATE_REPO
	container.SetGitPathPrivateRepository("github.com/__YOUR_GUTHUB__")

	// new image name delete:latest
	container.SetImageName("delete:latest")

	// container name container_delete_server_after_test
	container.SetContainerName("container_delete_server_after_test")

	// git project to clone git@github.com:helmutkemper/iotmaker.docker.builder.private.example.git
	container.SetGitCloneToBuild("git@github.com:__YOUR_GUTHUB__/__PROJECT__.git")

	// plase note, the main.go file must be at root folder
	container.MakeDefaultDockerfileForMe()

	// copy you ssh and git data to the container
	err = container.SetPrivateRepositoryAutoConfig()
	if err != nil {
		panic(err)
	}

	// set a waits for the text to appear in the standard container output to proceed [optional]
	container.SetWaitStringWithTimeout("Stating server on port 3000", 10*time.Second)

	// change and open port 3000 to 3030
	container.AddPortToChange("3000", "3030")

	// replace container folder /static to host folder ./test/static
	err = container.AddFiileOrFolderToLinkBetweenConputerHostAndContainer("./test/static", "/static")
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

	// build a new container from image
	err = container.ContainerBuildFromImage()
	if err != nil {
		panic(err)
	}
}
