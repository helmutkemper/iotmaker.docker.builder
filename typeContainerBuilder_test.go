package iotmaker_docker_builder

import (
	dockerNetwork "github.com/helmutkemper/iotmaker.docker.builder.network"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
	"log"
	"testing"
)

func TestContainer_1(t *testing.T) {
	var err error

	var dockerSys iotmakerdocker.DockerSystem
	err = dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = dockerSys.RemoveAllByNameContains("delete")
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	var netDocker = dockerNetwork.ContainerBuilderNetwork{}
	err = netDocker.Init()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = netDocker.NetworkCreate("cache_delete_after_test", "10.0.0.0/16", "10.0.0.1")
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	var container = ContainerBuilder{}
	container.SetNetworkDocker(&netDocker)
	container.SetImageName("nats:latest")
	container.SetContainerName("container_delete_nats_after_test")
	container.SetWaitString("Listening for route connections on 0.0.0.0:6222")

	err = container.Init()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = container.ImagePull()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = container.ContainerBuildFromImage()
	err = container.ImagePull()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = dockerSys.RemoveAllByNameContains("delete")
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}
}

func TestContainer_2(t *testing.T) {
	var err error

	var dockerSys iotmakerdocker.DockerSystem
	err = dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = dockerSys.RemoveAllByNameContains("delete")
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	var netDocker = dockerNetwork.ContainerBuilderNetwork{}
	err = netDocker.Init()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = netDocker.NetworkCreate("cache_delete_after_test", "10.0.0.0/16", "10.0.0.1")
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	var container = ContainerBuilder{}
	container.SetNetworkDocker(&netDocker)
	container.SetImageName("nats:latest")
	container.SetContainerName("container_delete_nats_after_test")
	container.AddPortToOpen("4222")
	container.SetWaitString("Listening for route connections on 0.0.0.0:6222")

	err = container.Init()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = container.ImagePull()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = container.ContainerBuildFromImage()
	err = container.ImagePull()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = dockerSys.RemoveAllByNameContains("delete")
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}
}

func TestContainer_3(t *testing.T) {
	var err error

	var dockerSys iotmakerdocker.DockerSystem
	err = dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = dockerSys.RemoveAllByNameContains("delete")
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	var netDocker = dockerNetwork.ContainerBuilderNetwork{}
	err = netDocker.Init()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = netDocker.NetworkCreate("cache_delete_after_test", "10.0.0.0/16", "10.0.0.1")
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	var container = ContainerBuilder{}
	container.SetNetworkDocker(&netDocker)
	container.SetImageName("nats:latest")
	container.SetContainerName("container_delete_nats_after_test")
	container.AddPortToChange("4222", "4200")
	container.SetWaitString("Listening for route connections on 0.0.0.0:6222")

	err = container.Init()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = container.ImagePull()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = container.ContainerBuildFromImage()
	err = container.ImagePull()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = dockerSys.RemoveAllByNameContains("delete")
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}
}

func TestContainer_4(t *testing.T) {
	var err error

	var dockerSys iotmakerdocker.DockerSystem
	err = dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = dockerSys.RemoveAllByNameContains("delete")
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	var netDocker = dockerNetwork.ContainerBuilderNetwork{}
	err = netDocker.Init()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = netDocker.NetworkCreate("cache_delete_after_test", "10.0.0.0/16", "10.0.0.1")
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	var container = ContainerBuilder{}
	container.SetNetworkDocker(&netDocker)
	container.SetImageName("delete:latest")
	container.SetContainerName("container_delete_nats_after_test")
	container.SetGitCloneToBuild("https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git")
	container.SetWaitString("Stating server on port 3000")

	err = container.Init()
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = container.ImageBuildFromServer()
	if err != nil {
		util.TraceToLog()
		log.Printf("container.ImageBuildFromServer().error: %v", err.Error())
		t.FailNow()
	}

	err = container.ContainerBuildFromImage()
	if err != nil {
		util.TraceToLog()
		log.Printf("container.ContainerBuildFromImage().error: %v", err.Error())
		t.FailNow()
	}

	err = dockerSys.RemoveAllByNameContains("delete")
	if err != nil {
		util.TraceToLog()
		log.Printf("container.RemoveAllByNameContains().error: %v", err.Error())
		t.FailNow()
	}
}
