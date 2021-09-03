package iotmakerdockerbuilder

func (e *ContainerBuilder) GetImageRepoDigests() (repoDigests []string) {
	return e.imageRepoDigests
}
