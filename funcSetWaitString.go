package iotmakerdockerbuilder

// SetWaitString
//
// English: Defines a text to be searched for in the container's default output and forces it to wait for the
// container to be considered ready-to-use
//
//   value: searched text
//
// Português: Define um texto a ser procurado na saída padrão do container e força a espera do mesmo para se
// considerar o container como pronto para uso
//
//   value: texto procurado
func (e *ContainerBuilder) SetWaitString(value string) {
	e.waitString = value
}
