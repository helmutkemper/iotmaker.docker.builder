package iotmakerdockerbuilder

func (e *ContainerBuilder) GetImageRepoDigests() (repoDigests []string) {

	if e.imageInspected == false {
		_, _ = e.ImageInspect()
	}

	return e.imageRepoDigests
}
