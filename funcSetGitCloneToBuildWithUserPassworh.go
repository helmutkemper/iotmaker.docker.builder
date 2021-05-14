package iotmaker_docker_builder

func (e *ContainerBuilder) SetGitCloneToBuildWithUserPassworh(url, user, password string) {
	e.gitData.url = url
	e.gitData.user = user
	e.gitData.password = password
}
