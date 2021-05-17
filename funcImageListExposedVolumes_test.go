package iotmaker_docker_builder

import (
	"fmt"
	"github.com/helmutkemper/util"
)

func ExampleContainerBuilder_ImageListExposedVolumes() {
	var err error
	var volumes []string

	GarbageCollector()

	var container = ContainerBuilder{}
	// new image name delete:latest
	container.SetImageName("delete:latest")
	// git project to clone https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git
	container.SetGitCloneToBuild("https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git")
	err = container.Init()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	err = container.ImageBuildFromServer()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	volumes, err = container.ImageListExposedVolumes()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	fmt.Printf("%v", volumes[0])

	GarbageCollector()

	// Output:
	// /static
}
