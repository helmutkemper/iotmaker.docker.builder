package iotmakerdockerbuilder

// AddImageBuildOptionsBuildArgs
//
// English: Set build-time variables (--build-arg)
//
//   key:   argument key (e.g. Dockerfile: ARG key)
//   value: argument value
//
//   docker build --build-arg HTTP_PROXY=http://10.20.30.2:1234
//   see https://docs.docker.com/engine/reference/commandline/build/#set-build-time-variables---build-arg
//
// Português: Adiciona uma variável durante a construção (--build-arg)
//
//   key:   chave do argumento (ex. Dockerfile: ARG key)
//   value: valor do argumento
//
//   docker build --build-arg HTTP_PROXY=http://10.20.30.2:1234
//   Veja https://docs.docker.com/engine/reference/commandline/build/#set-build-time-variables---build-arg
func (e *ContainerBuilder) AddImageBuildOptionsBuildArgs(key string, value *string) {
	if e.buildOptions.BuildArgs == nil {
		e.buildOptions.BuildArgs = make(map[string]*string)
	}

	e.buildOptions.BuildArgs[key] = value
}
