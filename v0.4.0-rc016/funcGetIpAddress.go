package iotmakerdockerbuilder

// GetIPV4Address (english):
//
// GetIPV4Address (português): Retorna o IP de quando o container foi criado.
//
//   Nota: Caso o container seja desconectado ou conectado a uma outra rede após a criação, esta informação pode mudar
//
func (e *ContainerBuilder) GetIPV4Address() (IP string) {
	return e.IPV4Address
}
