package iotmaker_docker_builder

func (e *ContainerBuilder) SetGitCloneToBuild(url string) {
	e.gitData.url = url
}
