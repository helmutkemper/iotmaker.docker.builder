package iotmaker_docker_builder

func (e *ContainerBuilder) SetGitCloneToBuildWithPrivateToken(url, privateToken string) {
	e.gitData.url = url
	e.gitData.privateToke = privateToken
}
