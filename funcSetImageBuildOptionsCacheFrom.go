package iotmakerdockerbuilder

// SetImageBuildOptionsCacheFrom (english): CacheFrom specifies images that are used for matching cache.
// Images specified here do not need to have a valid parent chain to match cache.
//
// SetImageBuildOptionsCacheFrom (português):
func (e *ContainerBuilder) SetImageBuildOptionsCacheFrom(values []string) {
	e.buildOptions.CacheFrom = values
}
