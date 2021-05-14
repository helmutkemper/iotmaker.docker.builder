package iotmaker_docker_builder

// gitData (Português): Estrutura de dados baseada no framework go-git
type gitData struct {
	url               string
	sshPrivateKeyPath string
	privateToke       string
	user              string
	password          string
}
