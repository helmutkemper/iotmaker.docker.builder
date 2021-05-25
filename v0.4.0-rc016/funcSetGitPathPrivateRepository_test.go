package iotmakerdockerbuilder

import (
	"fmt"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func ExampleContainerBuilder_SetGitPathPrivateRepository() {
	var err error

	GarbageCollector()

	var container = ContainerBuilder{}
	container.SetGitPathPrivateRepository("github.com/helmutkemper")
	// new image name delete:latest
	container.SetImageName("delete:latest")
	// container name container_delete_server_after_test
	container.SetContainerName("container_delete_server_after_test")
	// git project to clone git@github.com:helmutkemper/iotmaker.docker.builder.private.example.git
	container.SetGitCloneToBuild("git@github.com:helmutkemper/iotmaker.docker.builder.private.example.git")
	container.MakeDefaultDockerfileForMe()
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

	go func(ch *chan iotmakerdocker.ContainerPullStatusSendToChannel) {
		for {

			select {
			case event := <-*ch:
				var stream = event.Stream
				stream = strings.ReplaceAll(stream, "\n", "")
				stream = strings.ReplaceAll(stream, "\r", "")
				stream = strings.Trim(stream, " ")

				if stream == "" {
					continue
				}

				log.Printf("%v", stream)

				if event.Closed == true {
					return
				}
			}
		}
	}(container.GetChannelEvent())

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

	// At this point, the container is ready for use on port 3030

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
