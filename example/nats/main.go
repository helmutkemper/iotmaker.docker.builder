package main

import (
	builder "github.com/helmutkemper/iotmaker.docker.builder"
	dockerNetwork "github.com/helmutkemper/iotmaker.docker.builder.network"
	"time"
)

func main() {
	var err error

	var netDocker = dockerNetwork.ContainerBuilderNetwork{}
	err = netDocker.Init()
	if err != nil {
		panic(err)
	}

	// create a network named delete_after_test, subnet 10.0.0.0/16 e gatway 10.0.0.1
	err = netDocker.NetworkCreate("delete_after_test", "10.0.0.0/16", "10.0.0.1")
	if err != nil {
		panic(err)
	}

	// create a container
	var container = builder.ContainerBuilder{}
	// set image name for docker pull
	container.SetImageName("nats:latest")
	// link container and network [optional] (next ip address is 10.0.0.2)
	container.SetNetworkDocker(&netDocker)
	// set a container name
	container.SetContainerName("nats_delete_after_test")
	// set a waits for the text to appear in the standard container output to proceed [optional]
	container.SetWaitStringWithTimeout("Listening for route connections on 0.0.0.0:6222", 10*time.Second)

	// inialize the container object
	err = container.Init()
	if err != nil {
		panic(err)
	}

	// image nats:latest pull command [optional]
	err = container.ImagePull()
	if err != nil {
		panic(err)
	}

	// container build and start from image nats:latest
	// waits for the text "Listening for route connections on 0.0.0.0:6222" to appear  in the standard container
	// output to proceed
	err = container.ContainerBuildAndStartFromImage()
	if err != nil {
		panic(err)
	}

	// container nats_delete_after_test ready for use at this point of code
}
