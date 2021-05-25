package iotmakerdockerbuilder

// SetImageBuildOptionsSquash (english): squash the resulting image's layers to the parent
// preserves the original image and creates a new one from the parent with all the changes applied to a single layer
//
// SetImageBuildOptionsSquash (português):
func (e *ContainerBuilder) SetImageBuildOptionsSquash(value bool) {
	e.buildOptions.Squash = value
}
