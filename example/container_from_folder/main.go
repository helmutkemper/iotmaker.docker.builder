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
	// set a folder path to make a new image
	container.SetBuildFolderPath("./test/server")
	// container name container_delete_server_after_test
	container.SetContainerName("container_delete_server_after_test")
	// set a waits for the text to appear in the standard container output to proceed [optional]
	container.SetWaitStringWithTimeout("starting server at port 3000", 10*time.Second)
	// change and open port 3000 to 3030
	container.AddPortToExpose("3000")
	// replace container folder /static to host folder ./test/static
	err = container.AddFileOrFolderToLinkBetweenComputerHostAndContainer("../../test/static", "/static")
	if err != nil {
		panic(err)
	}

	// inicialize container object
	err = container.Init()
	if err != nil {
		panic(err)
	}

	// todo: fazer o inspect

	// builder new image from folder
	_, err = container.ImageBuildFromFolder()
	if err != nil {
		panic(err)
	}

	// build a new container from image
	err = container.ContainerBuildAndStartFromImage()
	if err != nil {
		panic(err)
	}
}
