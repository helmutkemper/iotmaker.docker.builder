package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"time"
)

func (e *ContainerBuilder) ImageInspect() (inspect types.ImageInspect, err error) {
	inspect, err = e.dockerSys.ImageInspect(e.imageID)
	if err != nil {
		util.TraceToLog()
		return
	}

	e.imageCreated, err = time.Parse(time.RFC3339Nano, inspect.Created)
	if err != nil {
		util.TraceToLog()
		return
	}

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
