package iotmakerdockerbuilder

// SetImageBuildOptionsCacheFrom
//
// English: specifies images that are used for matching cache.
// Images specified here do not need to have a valid parent chain to match cache.
//
// Português: especifica imagens que são usadas para correspondência de cache.
// As imagens especificadas aqui não precisam ter uma cadeia pai válida para corresponder a cache.
func (e *ContainerBuilder) SetImageBuildOptionsCacheFrom(values []string) {
	e.buildOptions.CacheFrom = values
}
