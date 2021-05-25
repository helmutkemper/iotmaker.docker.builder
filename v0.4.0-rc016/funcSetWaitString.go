package iotmakerdockerbuilder

// SetWaitString (english):
//
// SetWaitString (português): Define um texto a ser procurado na saída padrão do container e força a espera do
// mesmo para se considerar o container como pronto para uso
//   value: texto emitido na saída padrão informando por um evento esperado
func (e *ContainerBuilder) SetWaitString(value string) {
	e.waitString = value
}
