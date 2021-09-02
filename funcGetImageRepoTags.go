package iotmakerdockerbuilder

func (e *ContainerBuilder) GetImageRepoTags() (repoTags []string) {

	if e.imageInspected == false {
		_, _ = e.ImageInspect()
	}

	return e.imageRepoTags
}
