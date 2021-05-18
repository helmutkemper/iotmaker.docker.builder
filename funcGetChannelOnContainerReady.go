package iotmakerdockerbuilder

// GetChannelOnContainerReady (english):
//
// GetChannelOnContainerReady (português): Canal disparado quando o container está pronto para uso
//
//   Nota: Este canal espera o container sinalizar que está pronto, caso SetWaitString() não seja definido, ou espera
//   pelo texto definido em SetWaitString() aparecer na saída padrão
func (e *ContainerBuilder) GetChannelOnContainerReady() (channel *chan bool) {
	return e.onContainerReady
}
