package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"log"
)

// addImageBuildOptionsGitCredentials
//
// English: call from SetPrivateRepositoryAutoConfig()
//
// Português: chamada por SetPrivateRepositoryAutoConfig()
func (e *ContainerBuilder) addImageBuildOptionsGitCredentials() (buildOptions types.ImageBuildOptions) {

	if buildOptions.BuildArgs == nil {
		e.buildOptions.BuildArgs = make(map[string]*string)
	}

	if e.contentGitConfigFile != "" {
		e.buildOptions.BuildArgs["GITCONFIG_FILE"] = &e.contentGitConfigFile
		log.Printf("GITCONFIG_FILE: %v", e.contentGitConfigFile)
	}

	if e.contentKnownHostsFile != "" {
		e.buildOptions.BuildArgs["KNOWN_HOSTS_FILE"] = &e.contentKnownHostsFile
		log.Printf("KNOWN_HOSTS_FILE: %v", e.contentKnownHostsFile)
	}

	if e.contentIdRsaFile != "" {
		e.buildOptions.BuildArgs["SSH_ID_RSA_FILE"] = &e.contentIdRsaFile
		log.Printf("SSH_ID_RSA_FILE: %v", e.contentIdRsaFile)
	}

	if e.gitPathPrivateRepository != "" {
		e.buildOptions.BuildArgs["GIT_PRIVATE_REPO"] = &e.gitPathPrivateRepository
		log.Printf("GIT_PRIVATE_REPO: %v", e.gitPathPrivateRepository)
	}

	return
}
