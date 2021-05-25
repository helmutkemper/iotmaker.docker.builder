package iotmakerdockerbuilder

// SetContainerShellForShellFormOfRunCmdEntrypoint (english): shell for shell-form of run cmd entrypoint
//
// SetContainerShellForShellFormOfRunCmdEntrypoint (português): define o terminal (shell) para executar o entrypoint
func (e *ContainerBuilder) SetContainerShellForShellFormOfRunCmdEntrypoint(values []string) {
	e.containerConfig.Shell = values
}
