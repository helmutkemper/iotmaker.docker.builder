package iotmakerdockerbuilder

func (e *ContainerBuilder) GetGitCloneToBuild() (url string) {
	return e.gitData.url
}
