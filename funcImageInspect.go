package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"time"
)

func (e *ContainerBuilder) ImageInspect() (inspect types.ImageInspect, err error) {
	if e.imageID == "" {
		e.imageID, err = e.ImageFindIdByName(e.GetImageName())
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	inspect, err = e.dockerSys.ImageInspect(e.imageID)
	if err != nil {
		util.TraceToLog()
		return
	}
	log.Printf("inspect.Created: %v", inspect.Created)
	log.Printf("time.RFC3339: %v", time.RFC3339)
	e.imageCreated, err = time.Parse(time.RFC3339Nano, inspect.Created)
	if err != nil {
		log.Printf("error: %v", err.Error())
		util.TraceToLog()
		return
	}
	log.Printf("e.imageCreated: %v", e.imageCreated)
	e.imageInspected = true

	e.imageRepoTags = inspect.RepoTags
	e.imageRepoDigests = inspect.RepoDigests
	e.imageParent = inspect.Parent
	e.imageComment = inspect.Comment
	e.imageContainer = inspect.Container
	e.imageAuthor = inspect.Author
	e.imageArchitecture = inspect.Architecture
	e.imageVariant = inspect.Variant
	e.imageOs = inspect.Os
	e.imageOsVersion = inspect.OsVersion
	e.imageSize = inspect.Size
	e.imageVirtualSize = inspect.VirtualSize

	return
}
