package iotmaker_docker_builder

import (
	dockerNetwork "github.com/helmutkemper/iotmaker.docker.builder.network"
	"github.com/helmutkemper/util"
)

func ExampleContainerBuilder_ImagePull() {
	var err error

	// garbage collector delete all containers, images, volumes and networks whose name contains the term "delete"
	var garbageCollector = ContainerBuilder{}
	err = garbageCollector.Init()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	// set the term "delete" to garbage collector
	err = garbageCollector.RemoveAllByNameContains("delete")
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	// create a network [optional]
	var netDocker = dockerNetwork.ContainerBuilderNetwork{}
	err = netDocker.Init()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	// create a network named cache_delete_after_test, subnet 10.0.0.0/16 e gatway 10.0.0.1
	err = netDocker.NetworkCreate("cache_delete_after_test", "10.0.0.0/16", "10.0.0.1")
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	// create a container
	var container = ContainerBuilder{}
	// link container and network [optional] (next ip address is 10.0.0.2)
	container.SetNetworkDocker(&netDocker)
	// set image name for docker pull
	container.SetImageName("nats:latest")
	// set a container name
	container.SetContainerName("container_delete_nats_after_test")
	// set a waits for the text to appear in the standard container output to proceed [optional]
	container.SetWaitString("Listening for route connections on 0.0.0.0:6222")

	// inialize the container object
	err = container.Init()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	// image nats:latest pull command [optional]
	err = container.ImagePull()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	// container build and start from image nats:latest
	// waits for the text "Listening for route connections on 0.0.0.0:6222" to appear  in the standard container output
	// to proceed
	err = container.ContainerBuildFromImage()
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	// container "container_delete_nats_after_test" running and ready for use on this code point on IP 10.0.0.2
	// all nats ports are open
	// you can use AddPortToOpen("4222"), to open only ports defineds inside code;
	// you can use AddPortToChange("4222", "1111") to open only ports defineds inside code and change port 4222 to port
	// 1111;
	// you can use SetDoNotOpenContainersPorts() to not open containers ports

	// remove container and network after test
	err = garbageCollector.RemoveAllByNameContains("delete")
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	// use this function to remove image, ONLY before container stoped and deleted
	err = garbageCollector.ImageRemoveByName("nats:latest")
	if err != nil {
		util.TraceToLog()
		panic(err)
	}
}
