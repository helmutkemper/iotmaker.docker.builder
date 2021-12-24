package iotmakerdockerbuilder

// SetSceneNameOnChaosScene
//
// English:
//
//  Adds the container to a scene
//
//  Scenes help control the maximum amount of container stopped or paused at the same time
//
// Note:
//
//   * Use this function in conjunction with the ConfigChaosScene() function
//
// Português:
//
//  Adiciona o container a uma cena
//
//  Cenas ajudam a controlar a quantidade máxima de container parados ou pausados ao mesmo tempo
//
// Nota:
//
//   * Use esta função em conjunto com a função ConfigChaosScene()
func (e *ContainerBuilder) SetSceneNameOnChaosScene(name string) {
	e.chaos.sceneName = name
}
