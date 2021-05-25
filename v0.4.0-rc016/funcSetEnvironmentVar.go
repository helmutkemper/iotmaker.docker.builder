package iotmakerdockerbuilder

// SetEnvironmentVar (english):
//
// SetEnvironmentVar (português): Define as variáveis de ambiente
//   value: array de string contendo um variável de ambiente por chave
func (e *ContainerBuilder) SetEnvironmentVar(value []string) {
	e.environmentVar = value
}
