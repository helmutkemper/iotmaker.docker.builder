package iotmaker_docker_builder

import (
	"fmt"
	"github.com/helmutkemper/util"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ExampleContainerBuilder_AddPortToChange() {
	var err error

	GarbageCollector()

	var container = ContainerBuilder{}
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
	err = container.AddFiileOrFolderToLinkBetweenConputerHostAndContainer("./test/static", "/static")
	if err != nil {
		log.Printf("container.AddFiileOrFolderToLinkBetweenConputerHostAndContainer().error: %v", err.Error())
		util.TraceToLog()
		panic(err)
	}

	// inicialize container object
	err = container.Init()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	// builder new image from git project
	err = container.ImageBuildFromServer()
	if err != nil {
		util.TraceToLog()
		log.Printf("container.ImageBuildFromServer().error: %v", err.Error())
		panic(err)
	}

	// container build from image delete:latest
	err = container.ContainerBuildFromImage()
	if err != nil {
		util.TraceToLog()
		log.Printf("container.ContainerBuildFromImage().error: %v", err.Error())
		panic(err)
	}

	// container "container_delete_server_after_test" running and ready for use on this code point on port 3030

	// read server inside a container on address http://localhost:3030/
	var resp *http.Response
	resp, err = http.Get("http://localhost:3030/")
	if err != nil {
		util.TraceToLog()
		log.Printf("http.Get().error: %v", err.Error())
		panic(err)
	}

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		util.TraceToLog()
		log.Printf("http.Get().error: %v", err.Error())
		panic(err)
	}

	// print output
	fmt.Printf("%s", body)

	GarbageCollector()

	// Output:
	// <html><body><p>C is life! Golang is a evolution of C</p></body></html>
}
