package iotmaker_docker_builder

// GetLastLogs (english):
//
// GetLastLogs (português): Retorna a saída padrão do container baseado no último ciclo do ticker definido em
// SetInspectInterval()
//
//   Nota: a função GetChannelOnContainerInspect() retorna o canal disparado pelo ticker quando as informações estão
//   prontas para uso
func (e *ContainerBuilder) GetLastLogs() (logs string) {
	return e.logs
}
