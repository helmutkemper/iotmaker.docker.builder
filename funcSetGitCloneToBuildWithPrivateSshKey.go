package iotmaker_docker_builder

func (e *ContainerBuilder) SetGitCloneToBuildWithPrivateSshKey(url, privateSshKeyPath, password string) {
	e.gitData.url = url
	e.gitData.sshPrivateKeyPath = privateSshKeyPath
	e.gitData.password = password
}
