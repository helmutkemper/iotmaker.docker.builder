package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/mount"
	isolatedNetwork "github.com/helmutkemper/iotmaker.docker.builder.network.interface"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

// ContainerBuilder (english):
//
// ContainerBuilder (português): Gerenciador de containers e imagens docker
type ContainerBuilder struct {
	network            isolatedNetwork.ContainerBuilderNetworkInterface
	dockerSys          iotmakerdocker.DockerSystem
	changePointer      *chan iotmakerdocker.ContainerPullStatusSendToChannel
	onContainerReady   *chan bool
	onContainerInspect *chan bool
	imageName          string
	imageID            string
	containerName      string
	buildPath          string
	environmentVar     []string
	changePorts        []changePort
	openPorts          []string
	openAllPorts       bool
	waitString         string
	waitStringTimeout  time.Duration
	containerID        string
	ticker             *time.Ticker
	inspect            iotmakerdocker.ContainerInspect
	logs               string
	inspectInterval    time.Duration
	gitData            gitData
	volumes            []mount.Mount
	IPV4Address        string

	makeDefaultDockerfile bool

	contentIdRsaFile               string
	contentIdRsaFileWithScape      string
	contentKnownHostsFile          string
	contentKnownHostsFileWithScape string
	contentGitConfigFile           string
	contentGitConfigFileWithScape  string

	gitPathPrivateRepository string

	buildOptions types.ImageBuildOptions
}

func (e *ContainerBuilder) mountDefaultDockerfile() (dockerfile string, err error) {

	var info fs.FileInfo
	var found bool

	dockerfile += `
# (en) first stage of the process
# (pt) primeira etapa do processo
FROM golang:1.16-alpine as builder
#
`
	for k := range e.buildOptions.BuildArgs {
		switch k {
		case "SSH_ID_RSA_FILE":
			dockerfile += `
# (en) content from file /root/.ssh/id_rsa
# (pt) conteúdo do arquivo /root/.ssh/id_rsa
ARG SSH_ID_RSA_FILE
`
		case "KNOWN_HOSTS_FILE":
			dockerfile += `
# (en) content from file /root/.ssh/know_hosts
# (pt) conteúdo do arquivo /root/.ssh/know_hosts
ARG KNOWN_HOSTS_FILE
`
		case "GITCONFIG_FILE":
			dockerfile += `
# (en) content from file /root/.gitconfig
# (pt) conteúdo do arquivo /root/.gitconfig
ARG GITCONFIG_FILE
`
		case "GIT_PRIVATE_REPO":
			dockerfile += `
# (en) path from private repository. example: github.com/helmutkemper
# (pt) caminho do repositório privado. exemplo: github.com/helmutkemper
ARG GIT_PRIVATE_REPO
`
		default:
			dockerfile += `
ARG ` + k + `
`
		}
	}

	dockerfile += `
#
# (en) creates the .ssh directory within the root directory
# (pt) cria o diretório .ssh dentro do diretório root
RUN mkdir -p /root/.ssh/ && \
`
	_, found = e.buildOptions.BuildArgs["SSH_ID_RSA_FILE"]
	if found == true {
		dockerfile += `
    # (en) creates the id_esa file inside the .ssh directory
    # (pt) cria o arquivo id_esa dentro do diretório .ssh
    echo "$SSH_ID_RSA_FILE" > /root/.ssh/id_rsa && \
    # (en) adjust file access security
    # (pt) ajusta a segurança de acesso do arquivo
    chmod -R 600 /root/.ssh/ && \
`
	}

	_, found = e.buildOptions.BuildArgs["KNOWN_HOSTS_FILE"]
	if found == true {
		dockerfile += `
    # (en) creates the known_hosts file inside the .ssh directory
    # (pt) cria o arquivo known_hosts dentro do diretório .ssh
    echo "$KNOWN_HOSTS_FILE" > /root/.ssh/known_hosts && \
    # (en) adjust file access security
    # (pt) ajusta a segurança de acesso do arquivo
    chmod -R 600 /root/.ssh/known_hosts && \
`
	}

	_, found = e.buildOptions.BuildArgs["GITCONFIG_FILE"]
	if found == true {
		dockerfile += `
    # (en) creates the .gitconfig file at the root of the root directory
    # (pt) cria o arquivo .gitconfig na raiz do diretório /root
    echo "$GITCONFIG_FILE" > /root/.gitconfig && \
    # (en) adjust file access security
    # (pt) ajusta a segurança de acesso do arquivo
    chmod -R 600 /root/.gitconfig && \
`
	}

	dockerfile += `
    # (en) prepares the OS for installation
    # (pt) prepara o OS para instalação
    apk update && \
    # (en) install binutils, file, gcc, g++, make, libc-dev, fortify-headers and patch
    # (pt) instala binutils, file, gcc, g++, make, libc-dev, fortify-headers e patch
    apk add --no-cache build-base && \
    # (en) install git, fakeroot, scanelf, openssl, apk-tools, libc-utils, attr, tar, pkgconf, patch, lzip, curl, 
    #      /bin/sh, so:libc.musl-x86_64.so.1, so:libcrypto.so.1.1 and so:libz.so.1
    # (pt) instala git, fakeroot, scanelf, openssl, apk-tools, libc-utils, attr, tar, pkgconf, patch, lzip, curl, 
    #      /bin/sh, so:libc.musl-x86_64.so.1, so:libcrypto.so.1.1 e so:libz.so.1
    apk add --no-cache alpine-sdk && \
    # (en) clear the cache
    # (pt) limpa a cache
    rm -rf /var/cache/apk/*
#
# (en) creates the /app directory, where your code will be installed
# (pt) cria o diretório /app, onde seu código vai ser instalado
WORKDIR /app
# (en) copy your project into the /app folder
# (pt) copia seu projeto para dentro da pasta /app
COPY . .
# (en) enables the golang compiler to run on an extremely simple OS, scratch
# (pt) habilita o compilador do golang para rodar em um OS extremamente simples, o scratch
ARG CGO_ENABLED=0
# (en) adjust git to work with shh
# (pt) ajusta o git para funcionar com shh
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/
`

	_, found = e.buildOptions.BuildArgs["GIT_PRIVATE_REPO"]
	if found == true {
		dockerfile += `
# (en) defines the path of the private repository
# (pt) define o caminho do repositório privado
RUN echo "go env -w GOPRIVATE=$GIT_PRIVATE_REPO"
`
	}

	dockerfile += `
# (en) install the dependencies in the go.mod file
# (pt) instala as dependências no arquivo go.mod
RUN go mod tidy
# (en) compiles the main.go file
# (pt) compila o arquivo main.go
RUN go build -ldflags="-w -s" -o /app/main /app/main.go
# (en) creates a new scratch-based image
# (pt) cria uma nova imagem baseada no scratch
# (en) scratch is an extremely simple OS capable of generating very small images
# (pt) o scratch é um OS extremamente simples capaz de gerar imagens muito reduzidas
# (en) discarding the previous image erases git access credentials for your security and reduces the size of the 
#      image to save server space
# (pt) descartar a imagem anterior apaga as credenciais de acesso ao git para a sua segurança e reduz o tamanho 
#      da imagem para poupar espaço no servidor
FROM scratch
# (en) copy your project to the new image
# (pt) copia o seu projeto para a nova imagem
COPY --from=builder /app/main .
# (en) execute your project
# (pt) executa o seu projeto
`

	for _, v := range e.changePorts {
		dockerfile += `EXPOSE ` + v.oldPort + `
`
	}

	for _, v := range e.openPorts {
		dockerfile += `EXPOSE ` + v + `
`
	}

	for _, v := range e.volumes {
		if v.Type == iotmakerdocker.KVolumeMountTypeBindString {
			info, err = os.Stat(v.Source)
			if err != nil {
				return
			}
			if info.IsDir() == true {
				dockerfile += `VOLUME ` + v.Target + `
`
			} else {
				var dir string
				dir, _ = filepath.Split(v.Target)
				dockerfile += `VOLUME ` + dir + `
`
			}
		}

	}

	dockerfile += `
CMD ["/main"]
`

	strings.ReplaceAll(dockerfile, "\r", "")
	var lineList = strings.Split(dockerfile, "\n")
	dockerfile = ""
	for _, line := range lineList {
		if strings.Trim(line, " ") != "" {
			dockerfile += line + "\r\n"
		}
	}

	return
}

func (e *ContainerBuilder) MakeDefaultGolangDockerfileForMe() {
	e.makeDefaultDockerfile = true
}

// SetPrivateRepositoryAutoConfig (english):
//
// SetPrivateRepositoryAutoConfig (português): Copia a chave ssh  ~/.ssh/id_rsa e o arquivo ~/.gitconfig para as
// variáveis SSH_ID_RSA_FILE e GITCONFIG_FILE.
//
// Dentro do Dockerfile
func (e *ContainerBuilder) SetPrivateRepositoryAutoConfig() (err error) {
	var userData *user.User
	var fileData []byte

	userData, err = user.Current()
	if err != nil {
		return
	}

	var filePathToRead = filepath.Join(userData.HomeDir, ".ssh", "id_rsa")
	fileData, err = ioutil.ReadFile(filePathToRead)
	if err != nil {
		return
	}

	e.contentIdRsaFile = string(fileData)
	e.contentIdRsaFileWithScape = strings.ReplaceAll(e.contentIdRsaFile, `"`, `\"`)

	filePathToRead = filepath.Join(userData.HomeDir, ".ssh", "known_hosts")
	fileData, err = ioutil.ReadFile(filePathToRead)
	if err != nil {
		return
	}

	e.contentKnownHostsFile = string(fileData)
	e.contentKnownHostsFileWithScape = strings.ReplaceAll(e.contentKnownHostsFile, `"`, `\"`)

	filePathToRead = filepath.Join(userData.HomeDir, ".gitconfig")
	fileData, err = ioutil.ReadFile(filePathToRead)
	if err != nil {
		return
	}

	e.contentGitConfigFile = string(fileData)
	e.contentGitConfigFileWithScape = strings.ReplaceAll(e.contentGitConfigFile, `"`, `\"`)

	e.addImageBuildOptionsGitCredentials()
	return
}
