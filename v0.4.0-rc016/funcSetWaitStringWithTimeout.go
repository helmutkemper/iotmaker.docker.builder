package iotmakerdockerbuilder

import (
	"time"
)

// SetWaitStringWithTimeout (english):
//
// SetWaitStringWithTimeout (português): Define um texto a ser procurado na saída padrão do container e força a espera do
// mesmo para se considerar o container como pronto para uso
//   value: texto emitido na saída padrão informando por um evento esperado
//   timeout: tempo máximo de espera
func (e *ContainerBuilder) SetWaitStringWithTimeout(value string, timeout time.Duration) {
	e.waitString = value
	e.waitStringTimeout = timeout
}
