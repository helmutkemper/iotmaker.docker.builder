package iotmakerdockerbuilder

// SetImageBuildOptionsNoCache (english): Set image build no cache
//
// SetImageBuildOptionsNoCache (português):
func (e *ContainerBuilder) SetImageBuildOptionsNoCache(value bool) {
	e.buildOptions.NoCache = value
}
