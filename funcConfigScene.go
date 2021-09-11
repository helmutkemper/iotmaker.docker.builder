package iotmakerdockerbuilder

// ConfigScene
//
// English: Add and configure a test scene prevents all containers in the scene from stopping at the same time
//
// Português: Adiciona e configura uma cena de teste impede que todos os container da cena parem ao mesmo tempo
func ConfigScene(sceneName string, maxStopedContainers, maxPausedContainers int) {
	theater.ConfigScene(sceneName, maxStopedContainers, maxPausedContainers)
}
