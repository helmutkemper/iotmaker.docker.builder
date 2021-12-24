package iotmakerdockerbuilder

// ConfigChaosScene
//
// English:
//
//  Add and configure a test scene prevents all containers in the scene from stopping at the same time
//
//   Input:
//     sceneName: unique name for the scene
//     maxStopedContainers: Maximum number of stopped containers
//     maxPausedContainers: Maximum number of paused containers
//     maxTotalPausedAndStoppedContainers: Maximum number of containers stopped and paused at the same time
//
// Note:
//
//   * Use this function with SetSceneName() function.
//
// Português:
//
//  Adiciona e configura uma cena de teste impedindo que todos os container da cena parem ao mesmo tempo
//
//   Entrada:
//     sceneName: Nome único para a cena
//     maxStopedContainers: Quantidade máxima de containers parados
//     maxPausedContainers: Quantidade máxima de containers pausados
//     maxTotalPausedAndStoppedContainers: Quantidade máxima de containers parados e pausados ao mesmo tempo
//
// Nota:
//
//   * Use esta função em conjunto com a função SetSceneName().
func ConfigChaosScene(sceneName string, maxStopedContainers, maxPausedContainers, maxTotalPausedAndStoppedContainers int) {
	theater.ConfigScene(sceneName, maxStopedContainers, maxPausedContainers, maxTotalPausedAndStoppedContainers)
}
