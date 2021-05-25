package iotmakerdockerbuilder

// changePort (english):
//
// changePort (português): Recebe a relação entre portas a serem trocadas
//   oldPort: porta original da imagem
//   newPort: porta a exporta na rede
type changePort struct {
	oldPort string
	newPort string
}
