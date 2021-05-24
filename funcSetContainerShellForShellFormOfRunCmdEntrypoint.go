package iotmakerdockerbuilder

// SetContainerShellForShellFormOfRunCmdEntrypoint (english):
//
// SetContainerShellForShellFormOfRunCmdEntrypoint (português):
func (e *ContainerBuilder) SetContainerShellForShellFormOfRunCmdEntrypoint(values []string) {
	e.containerConfig.Shell = values
}
